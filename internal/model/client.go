package model

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	MAX_MESSAGE_SIZE         = 1048576
	REQUEST_TIMEOUT          = time.Second * 3
	RESPONSE_SEARCH_INTERVAL = time.Millisecond * 50
)

// IncomingHandler Обработчик входящих сообщений.
type IncomingHandler func(message []byte) error

// Client Клиент сервера chrome.
type Client struct {
	debug           *Debug
	con             *websocket.Conn
	wg              *sync.WaitGroup
	requests        chan []byte
	incomingHandler IncomingHandler
	done            chan struct{}
	Responses       map[int]*Response
	Mu              *sync.RWMutex
	requestId       int
}

// NewClient Конструктор.
func NewClient(d *Debug, i IncomingHandler) *Client {
	d.Debug("chrome client created")
	return &Client{
		debug:           d,
		wg:              &sync.WaitGroup{},
		requests:        make(chan []byte),
		incomingHandler: i,
		done:            make(chan struct{}),
		Responses:       make(map[int]*Response),
		Mu:              &sync.RWMutex{},
	}
}

// Open Установка соединения.
func (c *Client) Open(ctx context.Context, u string) error {
	var err error
	if c.con, _, err = websocket.DefaultDialer.DialContext(ctx, u, nil); err != nil {
		return err
	}
	c.debug.Debug("chrome server connection opened")
	c.wg.Add(1)
	go c.read()
	c.wg.Add(1)
	go c.write()
	return nil
}

// Exec Выполнить запрос и проигнорировать ответ.
func (c *Client) Exec(method string, params []byte, options ...RequestOption) error {
	_, err := c.Query(method, params, options...)
	if err != nil {
		return err
	}
	return nil
}

// Do Выполнить запрос и получить ответ.
func (c *Client) Query(method string, params []byte, options ...RequestOption) ([]byte, error) {
	c.requestId++
	req := &Request{
		Id:     c.requestId,
		Method: method,
		Params: params,
	}
	for _, o := range options {
		if err := o(req); err != nil {
			return nil, err
		}
	}
	c.requests <- req.Json()
	s := time.Now()
	for {
		if time.Now().UnixNano() > s.Add(REQUEST_TIMEOUT).UnixNano() {
			return nil, ERR_RESPONSE_TIMEOUT
		}
		c.debug.Debug(fmt.Sprintf("try catch response on request with id %d", req.Id))
		res, o := c.searchResponse(req.Id)
		if !o {
			time.Sleep(RESPONSE_SEARCH_INTERVAL)
			continue
		}
		c.Mu.Lock()
		delete(c.Responses, req.Id)
		c.Mu.Unlock()
		if res.Error != nil {
			return nil, c.responseError(req, res)
		}
		c.debug.Debug(fmt.Sprintf("response %d catched", req.Id))
		return res.Result, nil
	}
}

// Close Закрытие соединения.
func (c *Client) Close() error {
	if err := c.con.Close(); err != nil {
		return err
	}
	close(c.done)
	c.wg.Wait()
	return nil
}

// read Чтение.
func (c *Client) read() {
	defer func() {
		c.debug.Debug("reading messages stopped")
		c.wg.Done()
	}()
	c.con.SetReadLimit(MAX_MESSAGE_SIZE)
	c.debug.Debug("reading messages started")
	for {
		_, m, err := c.con.ReadMessage()
		if err != nil {
			if errors.Is(err, net.ErrClosed) ||
				websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				// соединение закрыто.
				return
			}
			fmt.Println("read err", err)
			return
		}
		c.incomingHandler(m)
	}
}

// write Запись.
func (c *Client) write() {
	defer func() {
		c.debug.Debug("writing messages stopped")
		c.wg.Done()
	}()
	c.debug.Debug("writing messages started")
	for {
		select {
		case m := <-c.requests:
			c.debug.Debug(fmt.Sprintf("-> %+v", string(m)))
			if err := c.con.WriteMessage(websocket.TextMessage, m); err != nil {
				fmt.Println("write err", err)
			}
		case <-c.done:
			return
		}
	}
}

// searchResponse Поиск ответа.
func (c *Client) searchResponse(id int) (*Response, bool) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	for rid, res := range c.Responses {
		if id == rid {
			return res, true
		}
	}
	return nil, false
}

// responseError Ошибка ответа.
func (c *Client) responseError(req *Request, res *Response) error {
	if len(res.Error.Data) > 0 {
		return fmt.Errorf(
			"%w code: %d; message: %s; data: %s",
			fmt.Errorf(
				"request error id: %d; method: %s",
				req.Id,
				req.Method,
			),
			res.Error.Code,
			res.Error.Message,
			res.Error.Data,
		)
	}
	return fmt.Errorf(
		"%w code: %d; message: %s",
		fmt.Errorf("%s error", req.Method),
		res.Error.Code,
		res.Error.Message,
	)
}

package tools

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"time"
)

const (
	TEST_SERVER_PORT = 8000
	CURL_PATH        = "/usr/bin/curl"
	TEST_INTERVAL    = time.Millisecond * 100
	TEST_TIMEOUT     = time.Second * 3
	TEMPLATE_PATH    = "./internal/tools/page.html"
)

var (
	TEST_TIMEOUT_ERROR = errors.New("WAIT TEST SERVER TIMEOUT")
)

// Server HTTP сервер для тестирования реализации методов Chrome devtools protocol.
var Server *http.Server

// TestServerStarted Тестовый сервер уже запущен.
func TestServerStarted() bool {
	_, err := exec.Command(CURL_PATH, fmt.Sprintf("http://localhost:%d/", TEST_SERVER_PORT)).Output()
	if err != nil {
		return false
	}
	return true
}

// WaitTestServer Ожижание тестового сервера.
func WaitTestServer() error {
	s := time.Now()
	for {
		if s.Add(TEST_TIMEOUT).UnixNano() < time.Now().UnixNano() {
			return TEST_TIMEOUT_ERROR
		}
		if !TestServerStarted() {
			time.Sleep(TEST_INTERVAL)
			continue
		}
		return nil
	}
}

// StartTestServer Запуск тестового сервера.
func StartTestServer() {
	if Server == nil {
		if err := createTestServer(); err != nil {
			panic(err)
		}
	}
	if err := Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

// StopTestServer Остановка тестовго сервера.
func StopTestServer() error {
	return Server.Shutdown(context.Background())
}

// createTestServer Создание сервера для проведения тестов.
func createTestServer() error {
	// шаблон
	t, err := template.ParseFiles(TEMPLATE_PATH)
	if err != nil {
		return err
	}
	// роутинг
	h := http.NewServeMux()
	h.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	})
	// запуск сервера
	Server = &http.Server{
		Addr:    fmt.Sprintf(":%d", TEST_SERVER_PORT),
		Handler: h,
	}
	return nil
}

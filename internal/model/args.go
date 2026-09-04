package model

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

const (
	ARG_NAME_LANG                     = "--lang"
	ARG_NAME_LOG_LEVEL                = "--log-level"
	ARG_NAME_USER_DATA_DIR            = "--user-data-dir"
	ARG_NAME_ABOUT_BLANK              = "about:blank"
	ARG_NAME_NO_STARTUP_WINDOW        = "--no-startup-window"
	ARG_NAME_NO_DEFAULT_BROWSER_CHECK = "--no-default-browser-check"
	ARG_NAME_NO_SANDBOX               = "--no-sandbox"
	ARG_NAME_TEST_TYPE                = "--test-type"
	ARG_NAME_DISABLE_DEV_SHM_USAGE    = "--disable-dev-shm-usage"
	ARG_NAME_NO_FIRST_RUN             = "--no-first-run"
	ARG_NAME_REMOTE_DEBUGGING_HOST    = "--remote-debugging-host"
	ARG_NAME_REMOTE_DEBUGGING_PORT    = "--remote-debugging-port"
	ARG_NAME_HEADLESS                 = "--headless"
	ARG_NAME_USER_AGENT               = "--user-agent"

	ARG_VALUE_LANG                  = "ru"
	ARG_VALUE_LOG_LEVEL             = 0
	ARG_VALUE_USER_DATA_DIR         = "chrome_user_data_dir"
	ARG_VALUE_REMOTE_DEBUGGING_HOST = "127.0.0.1"
	ARG_VALUE_REMOTE_DEBUGGING_PORT = 46060
)

// Args Аргументы.
type Args map[string]*string

// NewConfigArgs Конструктор.
func NewArgs() *Args {
	a := make(Args)
	a.Set(ARG_NAME_LANG, ARG_VALUE_LANG)
	a.Set(ARG_NAME_LOG_LEVEL, strconv.Itoa(ARG_VALUE_LOG_LEVEL))
	a.Set(ARG_NAME_USER_DATA_DIR, ARG_VALUE_USER_DATA_DIR)
	a.Set(ARG_NAME_NO_STARTUP_WINDOW)
	a.Set(ARG_NAME_NO_DEFAULT_BROWSER_CHECK)
	a.Set(ARG_NAME_NO_SANDBOX)
	a.Set(ARG_NAME_TEST_TYPE)
	a.Set(ARG_NAME_DISABLE_DEV_SHM_USAGE)
	a.Set(ARG_NAME_NO_FIRST_RUN)
	a.Set(ARG_NAME_REMOTE_DEBUGGING_HOST, ARG_VALUE_REMOTE_DEBUGGING_HOST)
	a.Set(ARG_NAME_REMOTE_DEBUGGING_PORT, strconv.Itoa(ARG_VALUE_REMOTE_DEBUGGING_PORT))
	a.Set(ARG_NAME_HEADLESS)
	return &a
}

// NewArgsFromFile Конструктор с получением данных из файла.
func NewArgsFromFile(filename string) (*Args, error) {
	a := make(Args)
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(f, &a); err != nil {
		return nil, err
	}
	return &a, nil
}

// Set Установка.
func (a *Args) Set(name string, value ...string) {
	if len(value) > 0 {
		(*a)[name] = &value[0]
		return
	}
	(*a)[name] = nil
}

// Get Получить значение аргумента.
func (a *Args) Get(name string) (string, bool) {
	v, o := (*a)[name]
	return *v, o
}

// Join Объединение имен и значений аргументов в срез строк.
func (a *Args) Join() []string {
	var s []string
	for n, v := range *a {
		if v != nil {
			s = append(s, fmt.Sprintf("%s=%s", n, *v))
			continue
		}
		s = append(s, fmt.Sprintf("%s", n))
	}
	return s
}

package main

import (
	"fmt"
	"log"

	"github.com/skvdmt/chrome"
	"github.com/skvdmt/chrome/internal/tools"
)

// main Точка входа в приложение.
func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

// start Запуск приложения.
func start() error {
	go tools.StartTestServer()
	if err := tools.WaitTestServer(); err != nil {
		return err
	}
	defer tools.StopTestServer()
	if err := headerTest(); err != nil {
		return err
	}
	return nil
}

// headerTest Тестирование заголовка.
func headerTest() error {
	fmt.Println("header test")
	// Создание драйвера.
	d, err := chrome.NewDriver(
		chrome.WithRemoveUserDataDirAfterClose(),
	)
	if err != nil {
		return err
	}
	// Открытие браузера.
	if err := d.Open(); err != nil {
		return err
	}
	defer func() {
		// Закрытие браузера.
		if err := d.Close(); err != nil {
			panic(err)
		}
	}()
	const (
		expected = "Welcome to Chrome devtools protocol testing page."
		selector = "#header"
	)
	// Перезод на страницу.
	if err := d.Navigate(fmt.Sprintf("http://localhost:%d/", tools.TEST_SERVER_PORT)); err != nil {
		return err
	}
	// Ожидание узла и получение его текстового содержимого.
	got, err := d.WaitNodeText(selector)
	if err != nil {
		return err
	}
	if got != expected {
		return fmt.Errorf(
			`error: node %s; text expetted: "%s"; got "%s"`,
			selector,
			expected,
			got,
		)
	}
	fmt.Println("OK")
	return nil
}

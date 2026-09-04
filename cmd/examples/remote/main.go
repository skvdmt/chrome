package main

import (
	"fmt"
	"log"

	"github.com/skvdmt/chrome"
)

// main Точка входа в приложение.
func main() {
	if err := headerTest(); err != nil {
		log.Fatal(err)
	}
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
		page     = "https://skvdmt.ru/"
		selector = "h1"
		expected = "Dmitry Skidanov"
	)
	// Перезод на страницу.
	if err := d.Navigate(page); err != nil {
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

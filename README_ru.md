# Golang Chrome DevTools protocol

## Переводы
[En](./README.md)

## Описание

Драйвер управления браузером.
Дравер реализует основные домены управления браузером Chrome DevTools Protocol:
- Browser - Управление браузером;
- DOM - Доступ к операциям чтения/записи объектной модели документа;
- Input - Ввод информации;
- Network - Сетевая активность страницы;
- Page - Взаимодействие со страницей;
- Target - Управление вкладками.

## Применение

- End-2-end тестирование.
- Сбор информации (scraping).
- Создание ботов.

## Установка

```
go get github.com/skvdmt/chrome
```

## Примеры
- Проверка заголовка
[локально](./cmd/examples/local/main.go) / [удаленно](./cmd/examples/remote/main.go)

## Ссылки

- [Источник](https://github.com/skvdmt/chrome)
- [Автор](https://skvdmt.ru/)
- [Документация управления Chrome](https://chromedevtools.github.io/devtools-protocol/)

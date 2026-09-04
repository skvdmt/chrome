# Golang Chrome DevTools protocol

## Translations
[Русский](./README_ru.md)

## Description
![Logo](./chrome.svg "Chrome")

Browser control driver.
Draver implements the core browser control domains of the Chrome DevTools Protocol:
- Browser - Browser control;
- DOM - Access to read/write operations for the Document Object Model;
- Input - Data entry;
- Network - Page network activity;
- Page - Interaction with the page;
- Target - Tab management.

## Used for

- End-to-end testing.
- Scraping.
- Creating bots.

## Instalation

```
go get github.com/skvdmt/chrome
```

## Examples
- Header check
[local](./cmd/examples/local/main.go) / [remote](./cmd/examples/remote/main.go)

## Links

- [Source](https://github.com/skvdmt/chrome)
- [Author](https://skvdmt.ru/)
- [Chrome management documentation](https://chromedevtools.github.io/devtools-protocol/)

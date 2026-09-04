package print_to_pdf

// Config Конфигурация.
type Config struct {
	// Ориентация бумаги. По умолчанию - false.
	Landscape bool `json:"landscape,omitempty"`
	// Отображать заголовок и нижний колонтитул. По умолчанию — false.
	DisplayHeaderFooter bool `json:"displayHeaderFooter,omitempty"`
	// Выводить фоновые изображения. По умолчанию — false.
	PrintBackground bool `json:"printBackground,omitempty"`
	// Масштаб отображения веб-страницы. По умолчанию — 1.
	Scale int `json:"scale,omitempty"`
	// Ширина бумаги указана в дюймах. По умолчанию — 8,5 дюймов.
	PaperWidth int `json:"paperWidth,omitempty"`
	// Высота листа бумаги указана в дюймах. По умолчанию — 11 дюймов.
	PaperHeight int `json:"paperHeight,omitempty"`
	// Верхнее поле в дюймах. По умолчанию — 1 см (~0,4 дюйма).
	MarginTop int `json:"marginTop,omitempty"`
	// Нижний отступ в дюймах. По умолчанию — 1 см (~0,4 дюйма).
	MarginBottom int `json:"marginBottom,omitempty"`
	// Левое поле в дюймах. По умолчанию — 1 см (~0,4 дюйма).
	MarginLeft int `json:"marginLeft,omitempty"`
	// Правое поле в дюймах. По умолчанию — 1 см (~0,4 дюйма).
	MarginRight int `json:"marginRight,omitempty"`
	// Укажите диапазоны страниц для печати, начиная с единицы, например,
	// «1-5, 8, 11-13». Страницы печатаются в порядке, указанном в документе,
	// а не в указанном порядке, и не более одного раза. По умолчанию
	// используется пустая строка, что означает печать всего документа.
	// Номера страниц ограничиваются фактическим количеством страниц документа,
	// а диапазоны за пределами конца документа игнорируются.
	// Если это приводит к отсутствию страниц для печати, сообщается об ошибке.
	// Указание диапазона с началом, превышающим конец, является ошибкой.
	PageRanges string `json:"pageRanges,omitempty"`
	// HTML-шаблон для заголовка при печати. ​​Должен представлять собой
	// допустимую HTML-разметку со следующими классами, используемыми
	// для внедрения значений при печати:
	//     date: отформатированная дата печати
	//     title: название документа
	//     url: местоположение документа
	//     pageNumber: текущий номер страницы
	//     totalPages: общее количество страниц в документе
	// Например, <span class=title></span> сгенерирует тег span, содержащий заголовок.
	HeaderTemplate string `json:"headerTemplate,omitempty"`
	// HTML-шаблон для нижнего колонтитула при печати.
	// Должен использовать тот же формат, что и заголовок шаблона.
	FooterTemplate string `json:"footerTemplate,omitempty"`
	// Определяет, следует ли отдавать предпочтение размеру страницы, заданному в CSS.
	// По умолчанию — false, в этом случае содержимое будет масштабировано
	// в соответствии с размером бумаги.
	PreferCSSPageSize bool `json:"preferCSSPageSize,omitempty"`
	// возвращать в виде потока
	// Допустимые значения: ReturnAsBase64, ReturnAsStream
	TransferMode string `json:"transferMode,omitempty"`
	// Создавать ли PDF-файл с тегами (доступный для просмотра).
	// По умолчанию используется выбор разработчика.
	GenerateTaggedPDF bool `json:"generateTaggedPDF,omitempty"`
	// Встраивать ли структуру документа в PDF-файл?
	GenerateDocumentOutline bool `json:"generateDocumentOutline,omitempty"`
}

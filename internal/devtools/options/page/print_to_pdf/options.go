package print_to_pdf

// Option Опция.
type Option func(c *Config)

// WithLandscape Ориентация бумаги.
func WithLandscape() Option {
	return func(c *Config) {
		c.Landscape = true
	}
}

// WithDisplayHeaderFooter Отображать заголовок и нижний колонтитул.
func WithDisplayHeaderFooter() Option {
	return func(c *Config) {
		c.DisplayHeaderFooter = true
	}
}

// WithPrintBackground Выводить фоновые изображения.
func WithPrintBackground() Option {
	return func(c *Config) {
		c.PrintBackground = true
	}
}

// WithScale Масштаб отображения веб-страницы.
func WithScale(scale int) Option {
	return func(c *Config) {
		c.Scale = scale
	}
}

// WithPaperWidth Ширина бумаги указана в дюймах.
func WithPaperWidth(paperWidth int) Option {
	return func(c *Config) {
		c.PaperWidth = paperWidth
	}
}

// WithPaperHeight Высота листа бумаги указана в дюймах.
func WithPaperHeight(paperHeight int) Option {
	return func(c *Config) {
		c.PaperHeight = paperHeight
	}
}

// WithMarginTop Верхнее поле в дюймах.
func WithMarginTop(marginTop int) Option {
	return func(c *Config) {
		c.MarginTop = marginTop
	}
}

// WithMarginBottom Нижний отступ в дюймах.
func WithMarginBottom(marginBottom int) Option {
	return func(c *Config) {
		c.MarginBottom = marginBottom
	}
}

// WithMarginLeft Левое поле в дюймах.
func WithMarginLeft(marginLeft int) Option {
	return func(c *Config) {
		c.MarginLeft = marginLeft
	}
}

// WithMarginRight Правое поле в дюймах.
func WithMarginRight(marginRight int) Option {
	return func(c *Config) {
		c.MarginRight = marginRight
	}
}

// WithPageRanges Диапазоны страниц для печати.
func WithPageRanges(pageRanges string) Option {
	return func(c *Config) {
		c.PageRanges = pageRanges
	}
}

// WithHeaderTemplate HTML-шаблон для заголовка при печати.
func WithHeaderTemplate(headerTemplate string) Option {
	return func(c *Config) {
		c.HeaderTemplate = headerTemplate
	}
}

// WithFooterTemplate HTML-шаблон для нижнего колонтитула при печати.
func WithFooterTemplate(footerTemplate string) Option {
	return func(c *Config) {
		c.FooterTemplate = footerTemplate
	}
}

// WithPreferCSSPageSize Отдавать предпочтение размеру страницы, заданному в CSS.
func WithPreferCSSPageSize() Option {
	return func(c *Config) {
		c.PreferCSSPageSize = true
	}
}

// WithTransferMode Возвращать в виде потока.
func WithTransferMode(transferMode string) Option {
	return func(c *Config) {
		c.TransferMode = transferMode
	}
}

// WithGenerateTaggedPDF Создавать PDF-файл с тегами.
func WithGenerateTaggedPDF() Option {
	return func(c *Config) {
		c.GenerateTaggedPDF = true
	}
}

// WithGenerateDocumentOutline Встраивать структуру документа в PDF-файл.
func WithGenerateDocumentOutline() Option {
	return func(c *Config) {
		c.GenerateDocumentOutline = true
	}
}

package dom

// BoxModel Box model.
type BoxModel struct {
	// Блок содержимого.
	Content *Quad `json:"content"`
	// Блок содержимого.
	Padding *Quad `json:"padding"`
	// Блок содержимого.
	Border *Quad `json:"border"`
	// Блок содержимого.
	Margin *Quad `json:"margin"`
	// Ширина узла.
	Width int `json:"width"`
	// Высота узла.
	Height int `json:"height"`
	// Форма вне координат.
	ShapeOutside *ShapeOutsideInfo `json:"shapeOutside"`
}

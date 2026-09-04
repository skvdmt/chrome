package dom

// ShapeOutsideInfo CSS Shape Внешние детали.
type ShapeOutsideInfo struct {
	Bounds      *Quad `json:"bounds"`
	Shape       []any `json:"shape"`
	MarginShape []any `json:"marginShape"`
}

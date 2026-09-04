package browser

// Bucket Ведро гистограммы Chrome.
type Bucket struct {
	// Минимальное значение (инклюзивное).
	Low int `json:"low"`
	// Максимальное значение (эксклюзивное).
	High int `json:"high"`
	// Количество образцов.
	Count int `json:"count"`
}

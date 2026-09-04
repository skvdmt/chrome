package browser

// Histogram Гистограмма Chrome.
type Histogram struct {
	// Имя.
	Name string `json:"name"`
	// Сумма значений примеров.
	Sum int `json:"sum"`
	// Общее количество примеров.
	Count int `json:"count"`
	// Ведра.
	Buckets []*Bucket `json:"buckets"`
}

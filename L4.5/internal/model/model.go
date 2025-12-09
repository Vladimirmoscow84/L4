package model

type Numbers struct {
	Data []float64 `json:"data"`
}

type Response struct {
	Sum    float64   `json:"sum"`
	Avg    float64   `json:"avg"`
	Median float64   `json:"median"`
	Sorted []float64 `json:"sorted"`
	Count  int       `json:"count"`
}

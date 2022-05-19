package first

type Item struct {
	Id     int     `json:"id,omitempty"`
	Treat  string  `json:"treat,omitempty"`
	Prewt  float64 `json:"prewt,omitempty"`
	Postwt float64 `json:"postwt,omitempty"`
}

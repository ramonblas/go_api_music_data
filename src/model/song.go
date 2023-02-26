package model

type Song struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Composer string `json:"director"`
}

type Songs []Song

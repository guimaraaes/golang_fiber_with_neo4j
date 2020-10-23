package model

type Movie struct {
	Title    string `json:"title"`
	Tagline  string `json:"tagline"`
	Released int64  `json:"released"`
}

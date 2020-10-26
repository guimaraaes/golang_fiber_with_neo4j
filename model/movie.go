package model

type Movie struct {
	// gorm.Model
	Title    string `json:"title"`
	Tagline  string `json:"tagline"`
	Released int64  `json:"released"`
}

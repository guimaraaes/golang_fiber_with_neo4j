package model

type Movie struct {
	title    string `json:"title"`
	tagline  string `json:"tagline"`
	released int    `json:"released"`
}

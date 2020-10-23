package model

type Person struct {
	TitleMovie   string `json:"titleMovie"`
	ReleaseMovie int    `json:"releaseMovie"`
	Name         string `json:"name"`
	Born         int    `json:"born"`
	Acted        bool   `json:"acted"`
	Direct       bool   `json:"direct"`
}

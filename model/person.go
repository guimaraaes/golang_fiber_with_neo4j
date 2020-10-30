package model

type Person struct {
	Name string `json:"name"`
	Born int    `json:"born"`
}

type KNOWS struct {
	// RelationType string `json:"relationship_type"`
	Ano string `json:"ano"`
	Mes string `json:"mes"`
}

type PersonKNOWSPerson struct {
	P1 Person
	KNOWS
	P2 Person
}

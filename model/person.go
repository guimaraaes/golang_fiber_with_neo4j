package model

type Person struct {
	Name string `json:"name"`
	Born int    `json:"born"`
}

type Relationship struct {
	RelationType string `json:"relationship_type"`
	// From         string `json:"from"`
	// To           string `json:"to"
	Ano string `json:"ano"`
	Mes string `json:"mes"`
}

// type RelationshipProperties struct {
// 	Ano string `json:"ano"`
// 	Mes string `json:"mes"`
// }

type PersonRelationship struct {
	Movie
	Person
	Relationship
}

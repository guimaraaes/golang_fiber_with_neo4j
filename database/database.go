package database

import (
	config "github.com/guimaraaes/golang_fiber_with_neo4j/config"

	neo4j "github.com/neo4j/neo4j-go-driver/neo4j"
)

var neo4j_Config = config.Neo4j{
	URI:      "bolt://35.153.83.239:32819",
	Username: "neo4j",
	Password: "default-cones-diseases",
	// URI:      "bolt://localhost:7687",
	// Username: "neo4j",
	// Password: "neo4j.",
}

type Neo4jDriverSession struct {
	Driver  neo4j.Driver
	Session neo4j.Session
}

var Neo4jDS Neo4jDriverSession

func ConnectionNeo4j() {
	d, err := neo4j.NewDriver(neo4j_Config.URI, neo4j.BasicAuth(neo4j_Config.Username, neo4j_Config.Password, ""))
	if err != nil {
		panic("failed to connect database")
	}
	// defer d.Close()
	s, err := d.Session(neo4j.AccessModeWrite)
	if err != nil {
		panic("failed to session database")
	}
	// defer s.Close()
	Neo4jDS.Driver = d
	Neo4jDS.Session = s
}

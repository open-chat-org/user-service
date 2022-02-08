package db

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

var driver neo4j.Driver

func InitNeo4j() {
	driver, _ = neo4j.NewDriver(
		"neo4j://localhost:7687",
		neo4j.BasicAuth("neo4j", "secret", ""),
	)
}

func GetNeo4j() neo4j.Driver {
	return driver
}

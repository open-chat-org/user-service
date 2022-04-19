package main

import (
	"user-service/controller"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	driver, err := neo4j.NewDriver(
		"neo4j://localhost:7687",
		neo4j.BasicAuth("neo4j", "secret", ""),
	)
	if err != nil {
		return
	}

	defer driver.Close()

	go controller.InitGrpc(driver)
	controller.InitGin(driver)
}

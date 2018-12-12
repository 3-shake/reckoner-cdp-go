package main

import (
	"fmt"
	"os"

	"github.com/3-shake/reckoner-cdp-go"
)

type Item struct {
	bar string
	foo string
}

func main() {
	accessKeyId := os.Getenv("RECKONER_CDP_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("RECKONER_CDP_SECRET_ACCESS_KEY")
	databaseName := "go_client_example_db"
	tableName := "go_client_example_table"

	client := reckonercdp.NewClient(accessKeyId, secretAccessKey, databaseName, tableName)

	if err := client.Insert(&Item{
		bar: "bar",
		foo: "foo",
	}); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("successfully uploaded!")
}

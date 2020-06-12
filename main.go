package main

import (
	"fmt"

	DB "example.com/m/src/db"
	"example.com/m/src/router"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Welcome to the server")

	e := router.New()

	DB.Connect()

	e.Start(":8000")

}

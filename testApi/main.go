package main

import (
	"mnc/testApi/db"
	"mnc/testApi/router"
)

func main() {

	db.Connect()

	router := router.SetupRouter()

	router.Run(":8080")
}

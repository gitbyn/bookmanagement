package main

import (
	"bookmanagement/database"
	"bookmanagement/routes"
)

func main() {
	database.Connect()
	r := routes.SetupRouter()
	r.Run(":8080")
}
package main

import (
	database "Mygram/databases"
	"Mygram/routers"
	"log"
)

func main() {
	database.InitDB()
	r := routers.InitApplication()
	log.Fatal(r.Run(":8080"))
}

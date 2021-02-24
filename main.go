package main

import (
	"log"

	"tinamar-api/api"
	"tinamar-api/config"
	"tinamar-api/persistence"
)

func main() {
	dbURI := config.MongoURL
	pers := persistence.Persistence{
		URI: dbURI,
	}

	conErr := pers.Connect()

	if conErr != nil {
		panic(conErr)
	}

	log.Println("Connected to database")

	restAPI := api.Server{
		Port: config.ServerPort,
		DB:   &pers,
	}

	restAPI.Init()
}

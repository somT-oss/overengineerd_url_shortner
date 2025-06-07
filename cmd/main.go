package main

import (
	"log"

	adapters "github.com/somT-oss/urlshortner/internals/adapters/db"
	"github.com/somT-oss/urlshortner/internals/application/api"
	"github.com/somT-oss/urlshortner/internals/application/urlshortner"
)

func main() {
	//postgres database url
	databaseUrl := ""

	// database adapter
	db, err := adapters.NewAdapter(databaseUrl)
	if err != nil {
		log.Fatal("the error %v occurred", err)
	}

	service, err := urlshortner.NewService()
	if err != nil {
		log.Fatalf("the error %v occurred", err)
	}

	application := api.NewApplication(
		db,
		service,
	)

	application.GenerateShortUrl("http://google.com")
}

package handlers

import (
	"fmt"
	"log"

	"olvido/pkg/mongo"
)

func Scraper() {
	mongoClient, err := mongo.Connect("root", "example", "", "")
	if err != nil {
		log.Fatalf("Failed to initiate mongoClient connection with given err: %s", err.Error())
	}
	for {
		// queries the mongodb for notifications to send
		fmt.Println("scrapping")
		fmt.Println(mongoClient)
	}
}

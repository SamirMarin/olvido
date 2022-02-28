package handlers

import (
	"fmt"
	"log"
	"time"

	"olvido/model"
	"olvido/pkg/mongo"
)

func Scraper() {
	mongoClient, err := mongo.Connect("root", "example", "", "", "olvido")
	if err != nil {
		log.Fatalf("Failed to initiate mongoClient connection with given err: %s", err.Error())
	}

	// this will be insered by the configuratin manager this is here for development purposes to be removed
	notifications := []model.Notification{
		model.Notification{
			Time:    time.Now().Unix(),
			Medium:  "email",
			Address: "marin.samir@gmail.com",
			Message: "Don't forget to study for your eks exam",
		},
		model.Notification{
			Time:    time.Now().Unix(),
			Medium:  "email",
			Address: "adrian.marin.estrada@gmail.com",
			Message: "Don't forget to study for your calc and line",
		},
		model.Notification{
			Time:    time.Now().Unix(),
			Medium:  "email",
			Address: "nair@gmail.com",
			Message: "Don't forget to study for your lsat",
		},
	}
	err = model.InsertNotifications(mongoClient, notifications)
	if err != nil {
		log.Fatalf("Failed to insert notifications with given err: %s", err.Error())
	}
	for {
		// queries the mongodb for notifications to send
		fmt.Println("scrapping")
		fmt.Println(mongoClient)
		//Todo make this configurable
		//get notifications from the db
		timeNow := time.Now()
		startTime := timeNow.Add(-(time.Hour * 24)).Unix()
		endTime := timeNow.Add(time.Hour * 24).Unix()
		notifications, err := model.GetNotificationsByTime(mongoClient, startTime, endTime)
		if err != nil {
			log.Fatalf("Failed to get notifications with given err: %s", err.Error())
		}
		//TODO this needs to be passed to the notifier so it can notify
		fmt.Println("These are the notifications from the DB")
		fmt.Println(notifications)
		time.Sleep(30 * time.Second)
	}
}

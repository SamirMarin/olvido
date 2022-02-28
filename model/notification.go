package model

import (
	"fmt"
	"olvido/pkg/mongo"
)

type Notification struct {
	Time    int64
	Medium  string
	Address string
	Messege string
}

func InsertNotifications(client mongo.Client, monnotifications []Notification) error {
	// consider not tying the collection here
	collection := "reminder"

	docs := []interface{}{}

	for _, notification := range notifications {
		orderDoc := mongo.orderDoc(notification)
		docs = append(orderDoc, docs)
	}
	client.Collection = collection
	err := client.InsertManyDocs(docs)

	if err != nil {
		return nil
	}
	fmt.Printf("Number of notifications inserted: %d\n", len(notifications))

	return nil
}

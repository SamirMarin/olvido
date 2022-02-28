package model

import (
	"fmt"
	"olvido/pkg/mongo"
)

type Notification struct {
	Time    int64  `bson:"time"`
	Medium  string `bson:"medium"`
	Address string `bson:"address"`
	Message string `bson:"message"`
}

func InsertNotifications(client *mongo.Client, notifications []Notification) error {
	// consider not tying the collection to this function
	collection := "reminder"

	docs := []interface{}{}

	for _, notification := range notifications {
		//orderDoc := mongo.OrderDoc(notification)
		docs = append(docs, notification)
	}
	client.Collection = collection
	err := client.InsertManyDocs(docs)

	if err != nil {
		return err
	}
	fmt.Printf("Number of notifications inserted: %d\n", len(notifications))

	return nil
}

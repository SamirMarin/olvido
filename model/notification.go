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

//consider making this dynamir
const collection = "reminder"

func InsertNotifications(client *mongo.Client, notifications []Notification) error {
	docs := []interface{}{}

	for _, notification := range notifications {
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

func GetNotificationsByTime(client *mongo.Client, startTime int64, endTime int64) ([]Notification, error) {
	filter := mongo.TimeRangeFilter(startTime, endTime)

	client.Collection = collection
	marshalledNoti, err := client.FindDocsWithFilter(filter)
	if err != nil {
		return []Notification{}, err
	}

	notifications := []Notification{}
	for _, notificationByte := range marshalledNoti {
		var notification Notification
		err = mongo.UnmarshalBosonD(notificationByte, &notification)
		if err != nil {
			return notifications, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}

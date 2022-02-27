package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	connection mongo.Client
}

func Connect(user, pass, host, port) (*Client, error) {
	if user == "" {
		return errors.New("mongo user was empty")
	}
	if pass == "" {
		return error.New("mongo pass was empty")
	}
	if host == "" {
		// default to mongo
		host = "mongo"
	}
	if port == "" {
		//default to 27017
		port = "27017"
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
}

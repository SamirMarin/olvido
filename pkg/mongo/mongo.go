package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	connection *mongo.Client
}

func Connect(user, pass, host, port string) (*Client, error) {
	if user == "" {
		return &Client{}, errors.New("mongo user was empty")
	}
	if pass == "" {
		return &Client{}, errors.New("mongo pass was empty")
	}
	if host == "" {
		// default to mongo
		host = "mongo"
	}
	if port == "" {
		//default to 27017
		port = "27017"
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority", user, pass, host, port)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return &Client{}, err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			//return &Client{}, err
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return &Client{}, err
	}
	fmt.Println(fmt.Sprintf("Successfully connected and pinged: %s.", uri))

	return &Client{
		connection: client,
	}, nil
}

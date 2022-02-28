package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	Connection *mongo.Client
	Database   string
	Collection string
}

func Connect(user, pass, host, port, database string) (*Client, error) {
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

	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority", user, pass, host, port)
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", user, pass, host, port)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return &Client{}, err
	}
	//TODO make sure we disconnect from db when done, needs to be moved to where connections are made
	//defer func() {
	//	if err = client.Disconnect(context.TODO()); err != nil {
	//		//return &Client{}, err
	//		panic(err)
	//	}
	//}()

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return &Client{}, err
	}
	fmt.Println(fmt.Sprintf("Successfully connected and pinged: %s.", uri))

	return &Client{
		Connection: client,
		Database:   database,
	}, nil
}

func OrderDoc(doc bson.D) interface{} {
	return bson.D(doc)
}

func (c *Client) InsertManyDocs(docs []interface{}) error {
	coll := c.Connection.Database(c.Database).Collection(c.Collection)

	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		return err
	}

	fmt.Printf("Number of documents inserted: %d\n", len(result.InsertedIDs))
	return nil
}

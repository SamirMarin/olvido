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

func (c *Client) InsertManyDocs(docs []interface{}) error {
	coll := c.Connection.Database(c.Database).Collection(c.Collection)

	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		return err
	}

	fmt.Printf("Number of documents inserted: %d\n", len(result.InsertedIDs))
	return nil
}

func TimeRangeFilter(start int64, end int64) interface{} {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"time", bson.D{{"$gt", start}}}},
				bson.D{{"time", bson.D{{"$lt", end}}}},
			}},
	}
	return filter
}

func (c *Client) FindDocsWithFilter(filter interface{}) ([][]byte, error) {
	coll := c.Connection.Database(c.Database).Collection(c.Collection)
	//TODO consider a projection decoupled form notifications mongo should not know about notifications
	//projection := bson.D{{"time", 1}, {"medium", 1}, {"message", 1}, {"_id", 0}}
	opts := options.Find() //.SetProjection(projection)
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		return [][]byte{}, err
	}

	var results []bson.D
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return [][]byte{}, err
	}

	marshalled := [][]byte{}

	for _, result := range results {
		doc, err := bson.Marshal(result)
		if err != nil {
			return [][]byte{}, err
		}
		marshalled = append(marshalled, doc)
	}

	return marshalled, nil
}

func UnmarshalBosonD(doc []byte, returnItem interface{}) error {
	err := bson.Unmarshal(doc, returnItem)
	if err != nil {
		return err
	}

	return nil
}

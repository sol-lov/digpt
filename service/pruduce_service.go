package service

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoService(uri, dbName, collectionName string) (*MongoService, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("error creating MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &MongoService{client: client, collection: collection}, nil
}

func (s *MongoService) SaveResponse(ctx context.Context, response []byte) error {
	session, err := s.client.StartSession()
	if err != nil {
		return fmt.Errorf("error starting MongoDB session: %v", err)
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		doc := bson.M{
			"response": response,
			"received": time.Now(),
		}

		_, err := s.collection.InsertOne(sessCtx, doc)
		if err != nil {
			return nil, fmt.Errorf("error inserting document into MongoDB: %v", err)
		}

		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return fmt.Errorf("error executing MongoDB transaction: %v", err)
	}

	return nil
}

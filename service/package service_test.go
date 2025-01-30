package service

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSaveResponse(t *testing.T) {
	// MongoDB connection URI
	uri := "mongodb://localhost:27017"

	// Create a new MongoService instance
	mongoService, err := NewMongoService(uri, "testdb", "testcollection")
	if err != nil {
		t.Fatalf("Failed to create MongoService: %v", err)
	}
	defer mongoService.client.Disconnect(context.Background())

	// Define the response to save
	response := []byte("test response")

	// Call SaveResponse method
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = mongoService.SaveResponse(ctx, response)
	if err != nil {
		t.Fatalf("Failed to save response: %v", err)
	}

	// Verify the document was inserted
	var result bson.M
	err = mongoService.collection.FindOne(ctx, bson.M{"response": response}).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to find inserted document: %v", err)
	}

	if string(result["response"].(primitive.Binary).Data) != string(response) {
		t.Errorf("Expected response %s, got %s", response, result["response"])
	}
}

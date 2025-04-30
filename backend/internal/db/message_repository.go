package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/ChewCEC/chat-app/internal/models"

)



func GetCollection(coll string) *mongo.Collection {

	err := InitMongo()
	if err != nil {
		log.Fatalf("Error initializing MongoDB: %v", err)
		
	}

	return MongoClient.Database("Chat-Information").Collection(coll)
}

func SaveMessage(message models.Message) error {
	// Set timestamp if not already set
	if message.Timestamp.IsZero() {
		message.Timestamp = time.Now()
	}

	collection := GetCollection("messages")
	_, err := collection.InsertOne(context.Background(), message)
	if err != nil {
		log.Printf("Error saving message: %v", err)
		return err
	}

	return nil
}

func GetRecentMessages(limit int64) ([]models.Message, error) {

	collection := GetCollection("messages")

	// Set options to sort by timestamp in descending order and limit results
	findOptions := options.Find().
		SetSort(bson.D{{Key: "timestamp", Value: -1}}).
		SetLimit(limit)

	cursor, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		log.Printf("Error retrieving messages: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var messages []models.Message
	if err = cursor.All(context.Background(), &messages); err != nil {
		log.Printf("Error decoding messages: %v", err)
		return nil, err
	}

	log.Printf("Retrieved %d messages", len(messages))

	// Reverse the order to get chronological order (oldest first)
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}

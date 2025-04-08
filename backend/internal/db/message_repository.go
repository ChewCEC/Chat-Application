package db

import (
	"context"
	"log"
	"time"

	"github.com/ChewCEC/chat-app/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	DATABASE   = "chat_app"
	COLLECTION = "messages"
)

// GetMessagesCollection returns the messages collection
func GetMessagesCollection() *mongo.Collection {
	return MongoClient.Database(DATABASE).Collection(COLLECTION)
}

// SaveMessage saves a message to the database
func SaveMessage(message models.Message) error {
	// Set timestamp if not already set
	if message.Timestamp.IsZero() {
		message.Timestamp = time.Now()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetMessagesCollection()
	_, err := collection.InsertOne(ctx, message)
	if err != nil {
		log.Printf("Error saving message: %v", err)
		return err
	}

	return nil
}

// GetRecentMessages retrieves the most recent messages from the database
func GetRecentMessages(limit int64) ([]models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetMessagesCollection()

	// Set options to sort by timestamp in descending order and limit results
	findOptions := options.Find().
		SetSort(bson.D{{Key: "timestamp", Value: -1}}).
		SetLimit(limit)

	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Printf("Error retrieving messages: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []models.Message
	if err = cursor.All(ctx, &messages); err != nil {
		log.Printf("Error decoding messages: %v", err)
		return nil, err
	}

	// Reverse the order to get chronological order (oldest first)
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}

package database

import (
	"context"
	"log"
	"time"

	"github.com/Ubivius/microservice-achievements/pkg/data"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAchievements struct {
	client     *mongo.Client
	collection *mongo.Collection
	logger     *log.Logger
}

func NewMongoAchievements(l *log.Logger) AchievementDB {
	mp := &MongoAchievements{logger: l}
	err := mp.Connect()
	// If connect fails, kill the program
	if err != nil {
		mp.logger.Fatal(err)
	}
	return mp
}

func (mp *MongoAchievements) Connect() error {
	// Setting client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:pass@localhost:27888/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil || client == nil {
		mp.logger.Fatalln("Failed to connect to database. Shutting down service")
	}

	// Ping DB
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		mp.logger.Fatal(err)
	}

	log.Println("Connection to MongoDB established")

	collection := client.Database("test").Collection("achievements")

	// Assign client and collection to the MongoAchievements struct
	mp.collection = collection
	mp.client = client
	return nil
}

func (mp *MongoAchievements) CloseDB() {
	err := mp.client.Disconnect(context.TODO())
	if err != nil {
		mp.logger.Println(err)
	} else {
		log.Println("Connection to MongoDB closed.")
	}
}

func (mp *MongoAchievements) GetAchievements() data.Achievements {
	// achievements will hold the array of Achievements
	var achievements data.Achievements

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Iterating through cursor
	for cursor.Next(context.TODO()) {
		var result data.Achievement
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		achievements = append(achievements, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return achievements
}

func (mp *MongoAchievements) GetAchievementByID(id string) (*data.Achievement, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.Achievement

	// Find a single matching item from the database
	err := mp.collection.FindOne(context.TODO(), filter).Decode(&result)

	// Parse result into the returned achievement
	return &result, err
}

func (mp *MongoAchievements) UpdateAchievement(achievement *data.Achievement) error {
	// Set updated timestamp in achievement
	achievement.UpdatedOn = time.Now().UTC().String()

	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: achievement.ID}}

	// Update sets the matched achievements in the database to achievement
	update := bson.M{"$set": achievement}

	// Update a single item in the database with the values in update that match the filter
	_, err := mp.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (mp *MongoAchievements) AddAchievement(achievement *data.Achievement) error {
	achievement.ID = uuid.NewString()
	// Adding time information to new achievement
	achievement.CreatedOn = time.Now().UTC().String()
	achievement.UpdatedOn = time.Now().UTC().String()

	// Inserting the new achievement into the database
	insertResult, err := mp.collection.InsertOne(context.TODO(), achievement)
	if err != nil {
		return err
	}

	log.Println("Inserting a document: ", insertResult.InsertedID)
	return nil
}

func (mp *MongoAchievements) DeleteAchievement(id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %v documents in the achievements collection\n", result.DeletedCount)
	return nil
}

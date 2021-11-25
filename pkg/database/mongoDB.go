package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Ubivius/microservice-achievements/pkg/data"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

// ErrorEnvVar : Environment variable error
var ErrorEnvVar = fmt.Errorf("missing environment variable")

type MongoAchievements struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoAchievements() AchievementDB {
	mp := &MongoAchievements{}
	err := mp.Connect()
	// If connect fails, kill the program
	if err != nil {
		log.Error(err, "MongoDB setup failed")
		os.Exit(1)
	}
	return mp
}

func (mp *MongoAchievements) Connect() error {
	uri := mongodbURI()

	// Setting client options
	opts := options.Client()
	clientOptions := opts.ApplyURI(uri)
	opts.Monitor = otelmongo.NewMonitor()

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Shutting down service")
		os.Exit(1)
	}

	// Ping DB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Error(err, "Failed to ping database. Shutting down service")
		os.Exit(1)
	}

	log.Info("Connection to MongoDB established")

	collection := client.Database("ubivius").Collection("achievements")

	// Assign client and collection to the MongoAchievements struct
	mp.collection = collection
	mp.client = client
	return nil
}

func (mp *MongoAchievements) PingDB() error {
	return mp.client.Ping(context.Background(), nil)
}

func (mp *MongoAchievements) CloseDB() {
	err := mp.client.Disconnect(context.Background())
	if err != nil {
		log.Error(err, "Error while disconnecting from database")
	}
}

func (mp *MongoAchievements) GetAchievements(ctx context.Context) data.Achievements {
	// achievements will hold the array of Achievements
	var achievements data.Achievements

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Error(err, "Error getting achievements from database")
	}

	// Iterating through cursor
	for cursor.Next(ctx) {
		var result data.Achievement
		err := cursor.Decode(&result)
		if err != nil {
			log.Error(err, "Error decoding achievement from database")
		}
		achievements = append(achievements, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err, "Error in cursor after iteration")
	}

	// Close the cursor once finished
	cursor.Close(ctx)

	return achievements
}

func (mp *MongoAchievements) GetAchievementByID(ctx context.Context, id string) (*data.Achievement, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.Achievement

	// Find a single matching item from the database
	err := mp.collection.FindOne(ctx, filter).Decode(&result)

	// Parse result into the returned achievement
	return &result, err
}

func (mp *MongoAchievements) UpdateAchievement(ctx context.Context, achievement *data.Achievement) error {
	// Set updated timestamp in achievement
	achievement.UpdatedOn = time.Now().UTC().String()

	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: achievement.ID}}

	// Update sets the matched achievements in the database to achievement
	update := bson.M{"$set": achievement}

	// Update a single item in the database with the values in update that match the filter
	_, err := mp.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error(err, "Error updating achievement.")
	}

	return err
}

func (mp *MongoAchievements) AddAchievement(ctx context.Context, achievement *data.Achievement) error {
	achievement.ID = uuid.NewString()
	// Adding time information to new achievement
	achievement.CreatedOn = time.Now().UTC().String()
	achievement.UpdatedOn = time.Now().UTC().String()

	// Inserting the new achievement into the database
	insertResult, err := mp.collection.InsertOne(ctx, achievement)
	if err != nil {
		return err
	}

	log.Info("Inserting achievement", "Inserted ID", insertResult.InsertedID)
	return nil
}

func (mp *MongoAchievements) DeleteAchievement(ctx context.Context, id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err, "Error deleting achievement")
	}

	log.Info("Deleted documents in achievements collection", "delete_count", result.DeletedCount)
	return nil
}

func deleteAllAchievementsFromMongoDB() error {
	uri := mongodbURI()

	// Setting client options
	opts := options.Client()
	clientOptions := opts.ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Failing test")
		return err
	}
	collection := client.Database("ubivius").Collection("achievements")
	_, err = collection.DeleteMany(context.Background(), bson.D{{}})
	return err
}

func mongodbURI() string {
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	if hostname == "" || port == "" || username == "" || password == "" {
		log.Error(ErrorEnvVar, "Some environment variables are not available for the DB connection. DB_HOSTNAME, DB_PORT, DB_USERNAME, DB_PASSWORD")
		os.Exit(1)
	}

	return "mongodb://" + username + ":" + password + "@" + hostname + ":" + port + "/?authSource=admin"
}

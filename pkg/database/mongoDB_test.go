package database

import (
	"context"
	"testing"

	"github.com/Ubivius/microservice-achievements/pkg/data"
	"github.com/google/uuid"
)

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoAchievements()
	if mp == nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBAddAchievementIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	achievement := &data.Achievement{
		Name:        "testName",
		Description: "testDescription",
		Condition:   "testCondition",
		SpriteID:    uuid.NewString(),
	}

	mp := NewMongoAchievements()
	err := mp.AddAchievement(context.Background(), achievement)
	if err != nil {
		t.Errorf("Failed to add achievement to database")
	}
	mp.CloseDB()
}

func TestMongoDBUpdateAchievementIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	achievement := &data.Achievement{
		ID:          uuid.NewString(),
		Name:        "testName",
		Description: "testDescription",
		Condition:   "testCondition",
		SpriteID:    uuid.NewString(),
	}

	mp := NewMongoAchievements()
	err := mp.UpdateAchievement(context.Background(), achievement)
	if err != nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBGetAchievementsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoAchievements()
	achievements := mp.GetAchievements(context.Background())
	if achievements == nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetAchievementByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoAchievements()
	_, err := mp.GetAchievementByID(context.Background(), "c9ddfb2f-fc4d-40f3-87c0-f6713024a993")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}

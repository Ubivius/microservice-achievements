package database

import (
	"context"
	"os"
	"testing"

	"github.com/Ubivius/microservice-achievements/pkg/data"
	"github.com/google/uuid"
)

func integrationTestSetup(t *testing.T) {
	t.Log("Test setup")

	if os.Getenv("DB_USERNAME") == "" {
		os.Setenv("DB_USERNAME", "admin")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		os.Setenv("DB_PASSWORD", "pass")
	}
	if os.Getenv("DB_PORT") == "" {
		os.Setenv("DB_PORT", "27888")
	}
	if os.Getenv("DB_HOSTNAME") == "" {
		os.Setenv("DB_HOSTNAME", "localhost")
	}

	err := deleteAllAchievementsFromMongoDB()
	if err != nil {
		t.Errorf("Failed to delete existing items from database during setup")
	}
}

func addAchievementAndGetId(t *testing.T) string {
	t.Log("Adding product")
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

	t.Log("Fetching new achievement ID")
	achievements := mp.GetAchievements(context.Background())
	mp.CloseDB()
	return achievements[len(achievements)-1].ID
}

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

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
	integrationTestSetup(t)

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

	achievements := mp.GetAchievements(context.Background())
	if achievements == nil || len(achievements) != 1 {
		t.Error("Incorrect number of returned achivements")
	}
	if achievements != nil && achievements[0].Name != achievement.Name {
		t.Errorf("Achievement is not the same. Expected name : %s but got %s", achievement.Name, achievements[0].Name)
	}
	mp.CloseDB()
}

func TestMongoDBUpdateAchievementIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	achievementID := addAchievementAndGetId(t)

	achievement := &data.Achievement{
		ID:          achievementID,
		Name:        "newName",
		Description: "testDescription",
		Condition:   "testCondition",
		SpriteID:    uuid.NewString(),
	}

	mp := NewMongoAchievements()
	err := mp.UpdateAchievement(context.Background(), achievement)
	if err != nil {
		t.Fail()
	}

	updatedAchievement, err := mp.GetAchievementByID(context.Background(), achievementID)
	if err != nil {
		t.Error("Error getting achievement from database")
	}
	if updatedAchievement.Name != achievement.Name {
		t.Errorf("Update on achievement failed, expected %s but got %s", achievement.Name, updatedAchievement.Name)
	}

	mp.CloseDB()
}

func TestMongoDBGetAchievementsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	addAchievementAndGetId(t)

	mp := NewMongoAchievements()
	achievements := mp.GetAchievements(context.Background())
	if achievements == nil {
		t.Error("Returned achievements is nil")
	}

	if len(achievements) != 1 {
		t.Errorf("Too many achievements in database, expected 1 but got %d", len(achievements))
	}
	mp.CloseDB()
}

func TestMongoDBGetAchievementByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	achievementID := addAchievementAndGetId(t)

	mp := NewMongoAchievements()
	_, err := mp.GetAchievementByID(context.Background(), achievementID)
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}

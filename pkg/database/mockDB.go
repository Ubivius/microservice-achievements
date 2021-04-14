package database

import (
	"time"

	"github.com/Ubivius/microservice-achievements/pkg/data"
	"github.com/google/uuid"
)

type MockAchievements struct {
}

func NewMockAchievements() AchievementDB {
	log.Info("Connecting to mock database")
	return &MockAchievements{}
}

func (mp *MockAchievements) Connect() error {
	return nil
}

func (mp *MockAchievements) PingDB() error {
	return nil
}

func (mp *MockAchievements) CloseDB() {
	log.Info("Mocked DB connection closed")
}

func (mp *MockAchievements) GetAchievements() data.Achievements {
	return achievementList
}

func (mp *MockAchievements) GetAchievementByID(id string) (*data.Achievement, error) {
	index := findIndexByAchievementID(id)
	if index == -1 {
		return nil, data.ErrorAchievementNotFound
	}
	return achievementList[index], nil
}

func (mp *MockAchievements) UpdateAchievement(achievement *data.Achievement) error {
	index := findIndexByAchievementID(achievement.ID)
	if index == -1 {
		return data.ErrorAchievementNotFound
	}
	achievementList[index] = achievement
	return nil
}

func (mp *MockAchievements) AddAchievement(achievement *data.Achievement) error {
	achievement.ID = uuid.NewString()
	achievementList = append(achievementList, achievement)
	return nil
}

func (mp *MockAchievements) DeleteAchievement(id string) error {
	index := findIndexByAchievementID(id)
	if index == -1 {
		return data.ErrorAchievementNotFound
	}

	achievementList = append(achievementList[:index], achievementList[index+1:]...)

	return nil
}

// Returns the index of a achievement in the database
// Returns -1 when no achievement is found
func findIndexByAchievementID(id string) int {
	for index, achievement := range achievementList {
		if achievement.ID == id {
			return index
		}
	}
	return -1
}

////////////////////////////////////////////////////////////////////////////////
/////////////////////////// Mocked database ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

var achievementList = []*data.Achievement{
	{
		ID:          "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Name:        "10 wins",
		Description: "Easy Peasy Lemon Squeezy",
		Condition:   "Accumulate 10 wins",
		SpriteID:    "a2181017-5c53-422b-b6bc-036b27c04fc8",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		Name:        "100 kills",
		Description: "Monster hunter",
		Condition:   "Accumulate 100 kills",
		SpriteID:    "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

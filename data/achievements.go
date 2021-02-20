package data

import (
	"fmt"
	"time"
)

// ErrorAchievementNotFound : Achievement specific errors
var ErrorAchievementNotFound = fmt.Errorf("Achievement not found")

// Achievement defines the structure for an API achievement.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Achievement struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Condition   string `json:"condition" validate:"required"`
	SpriteID    int    `json:"spriteid"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

// Achievements is a collection of Achievement
type Achievements []*Achievement

// All of these functions will become database calls in the future
// GETTING ACHIEVEMENTS

// GetAchievements : Returns the list of achievements
func GetAchievements() Achievements {
	return achievementList
}

// GetAchievementByID : Returns a single achievement with the given id
func GetAchievementByID(id int) (*Achievement, error) {
	index := findIndexByAchievementID(id)
	if id == -1 {
		return nil, ErrorAchievementNotFound
	}
	return achievementList[index], nil
}

// UPDATING ACHIEVEMENTS

// UpdateAchievement : need to remove id int from parameters when achievement handler is updated
func UpdateAchievement(achievement *Achievement) error {
	index := findIndexByAchievementID(achievement.ID)
	if index == -1 {
		return ErrorAchievementNotFound
	}
	achievementList[index] = achievement
	return nil
}

// AddAchievement : ADD A ACHIEVEMENT
func AddAchievement(achievement *Achievement) {
	achievement.ID = getNextID()
	achievementList = append(achievementList, achievement)
}

// DeleteAchievement : DELETING A ACHIEVEMENT
func DeleteAchievement(id int) error {
	index := findIndexByAchievementID(id)
	if index == -1 {
		return ErrorAchievementNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	achievementList = append(achievementList[:index], achievementList[index+1])

	return nil
}

// Returns the index of a achievement in the database
// Returns -1 when no achievement is found
func findIndexByAchievementID(id int) int {
	for index, achievement := range achievementList {
		if achievement.ID == id {
			return index
		}
	}
	return -1
}

//////////////////////////////////////////////////////////////////////////////
/////////////////////////// Fake database ///////////////////////////////////
///// DB connection setup and docker file will be done in sprint 8 /////////
///////////////////////////////////////////////////////////////////////////

// Finds the maximum index of our fake database and adds 1
func getNextID() int {
	lastAchievement := achievementList[len(achievementList)-1]
	return lastAchievement.ID + 1
}

// achievementList is a hard coded list of achievements for this
// example data source. Should be replaced by database connection
var achievementList = []*Achievement{
	{
		ID:          1,
		Name:        "10 wins",
		Description: "Easy Peasy Lemon Squeezy",
		Condition:   "Accumulate 10 wins",
		SpriteID:    14,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "100 kills",
		Description: "Monster hunter",
		Condition:   "Accumulate 100 kills",
		SpriteID:    69,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

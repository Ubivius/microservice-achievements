package data

import (
	"fmt"
)

// ErrorAchievementNotFound : Achievement specific errors
var ErrorAchievementNotFound = fmt.Errorf("Achievement not found")

// Achievement defines the structure for an API achievement.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Achievement struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Condition   string `json:"condition" validate:"required"`
	SpriteID    int    `json:"spriteid"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
}

// Achievements is a collection of Achievement
type Achievements []*Achievement

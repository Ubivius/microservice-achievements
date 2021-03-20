package database

import (
	"github.com/Ubivius/microservice-achievements/pkg/data"
)

// The interface that any kind of database must implement
type AchievementDB interface {
	GetAchievements() data.Achievements
	GetAchievementByID(id string) (*data.Achievement, error)
	UpdateAchievement(achievement *data.Achievement) error
	AddAchievement(achievement *data.Achievement) error
	DeleteAchievement(id string) error
	Connect() error
	CloseDB()
}

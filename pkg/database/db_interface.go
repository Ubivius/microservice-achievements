package database

import (
	"context"

	"github.com/Ubivius/microservice-achievements/pkg/data"
)

// The interface that any kind of database must implement
type AchievementDB interface {
	GetAchievements(ctx context.Context) data.Achievements
	GetAchievementByID(ctx context.Context, id string) (*data.Achievement, error)
	UpdateAchievement(ctx context.Context, achievement *data.Achievement) error
	AddAchievement(ctx context.Context, achievement *data.Achievement) error
	DeleteAchievement(ctx context.Context, id string) error
	Connect() error
	PingDB() error
	CloseDB()
}

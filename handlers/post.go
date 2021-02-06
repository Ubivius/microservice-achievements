package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/data"
)

// /POST /achievements
// Creates a new achievement
func (achievementHandler *AchievementsHandler) AddAchievement(responseWriter http.ResponseWriter, request *http.Request) {
	achievementHandler.logger.Println("Handle POST Achievement")
	achievement := request.Context().Value(KeyAchievement{}).(*data.Achievement)

	data.AddAchievement(achievement)
}
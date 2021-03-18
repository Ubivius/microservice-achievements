package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/data"
)

// AddAchievement creates a new achievement from the received JSON
func (achievementHandler *AchievementsHandler) AddAchievement(responseWriter http.ResponseWriter, request *http.Request) {
	achievementHandler.logger.Println("Handle POST Achievement")
	achievement := request.Context().Value(KeyAchievement{}).(*data.Achievement)

	data.AddAchievement(achievement)
	responseWriter.WriteHeader(http.StatusNoContent)
}

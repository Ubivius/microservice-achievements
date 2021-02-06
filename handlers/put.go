package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/data"
)

func (achievementHandler *AchievementsHandler) UpdateAchievements(responseWriter http.ResponseWriter, request *http.Request) {
	achievement := request.Context().Value(KeyAchievement{}).(data.Achievement)
	achievementHandler.logger.Println("Handle PUT achievement", achievement.ID)

	// Update achievement
	err := data.UpdateAchievement(&achievement)
	if err == data.ErrorAchievementNotFound {
		achievementHandler.logger.Println("[ERROR} achievement not found", err)
		http.Error(responseWriter, "Achievement not found", http.StatusNotFound)
		return
	}

	// Returns status, no content required
	responseWriter.WriteHeader(http.StatusNoContent)
}
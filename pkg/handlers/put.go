package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/data"
)

// UpdateAchievements updates the achievement with the ID specified in the received JSON achievement
func (achievementHandler *AchievementsHandler) UpdateAchievements(responseWriter http.ResponseWriter, request *http.Request) {
	achievement := request.Context().Value(KeyAchievement{}).(*data.Achievement)
	log.Info("UpdateAchievements request", "id", achievement.ID)

	// Update achievement
	err := achievementHandler.db.UpdateAchievement(achievement)
	if err == data.ErrorAchievementNotFound {
		log.Error(err, "Achievement not found")
		http.Error(responseWriter, "Achievement not found", http.StatusNotFound)
		return
	}

	// Returns status, no content required
	responseWriter.WriteHeader(http.StatusNoContent)
}

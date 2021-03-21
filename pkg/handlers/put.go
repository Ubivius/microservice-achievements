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
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
	case data.ErrorAchievementNotFound:
		log.Error(err, "Achievement not found")
		http.Error(responseWriter, "Achievement not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error updating achievement")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

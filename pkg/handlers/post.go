package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/data"
)

// AddAchievement creates a new achievement from the received JSON
func (achievementHandler *AchievementsHandler) AddAchievement(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("AddAchievement request")
	achievement := request.Context().Value(KeyAchievement{}).(*data.Achievement)

	err := achievementHandler.db.AddAchievement(achievement)

	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
	default:
		log.Error(err, "Error adding achievement")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

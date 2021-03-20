package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/data"
)

// Delete a achievement with specified id from the database
func (achievementHandler *AchievementsHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getAchievementID(request)
	log.Info("Delete achievement by ID request", "id", id)

	err := achievementHandler.db.DeleteAchievement(id)
	if err == data.ErrorAchievementNotFound {
		log.Error(err, "Error deleting achievement, id does not exist")
		http.Error(responseWriter, "Achievement not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Error(err, "Error deleting achievement")
		http.Error(responseWriter, "Error deleting achievement", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}

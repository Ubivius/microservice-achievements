package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/data"
)

// Delete a achievement with specified id from the database
func (achievementHandler *AchievementsHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getAchievementID(request)
	achievementHandler.logger.Println("Handle DELETE achievement", id)

	err := data.DeleteAchievement(id)
	if err == data.ErrorAchievementNotFound {
		achievementHandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Achievement not found", http.StatusNotFound)
		return
	}

	if err != nil {
		achievementHandler.logger.Println("[ERROR] deleting achievement", err)
		http.Error(responseWriter, "Error deleting achievement", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}

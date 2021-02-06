package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/data"
)

// GET /achievements
// Returns the full list of achievements
func (achievementHandler *AchievementsHandler) GetAchievements(responseWriter http.ResponseWriter, request *http.Request) {
	achievementHandler.logger.Println("Handle GET achievements")
	achievementList := data.GetAchievements()
	err := data.ToJSON(achievementList, responseWriter)
	if err != nil {
		achievementHandler.logger.Println("[ERROR] serializing achievement", err)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GET /achievements/{id}
// Returns a single achievement from the database
func (achievementHandler *AchievementsHandler) GetAchievementById(responseWriter http.ResponseWriter, request *http.Request) {
	id := getAchievementId(request)

	achievementHandler.logger.Println("[DEBUG] getting id", id)

	achievement, err := data.GetAchievementById(id)
	switch err {
	case nil:
	case data.ErrorAchievementNotFound:
		achievementHandler.logger.Println("[ERROR] fetching achievement", err)
		http.Error(responseWriter, "Achievement not found", http.StatusBadRequest)
		return
	default:
		achievementHandler.logger.Println("[ERROR] fetching achievement", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(achievement, responseWriter)
	if err != nil {
		achievementHandler.logger.Println("[ERROR] serializing achievement", err)
	}
}
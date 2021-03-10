package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/data"
)

// GetAchievements returns the full list of achievements
func (achievementHandler *AchievementsHandler) GetAchievements(responseWriter http.ResponseWriter, request *http.Request) {
	achievementHandler.logger.Println("Handle GET achievements")
	achievementList := data.GetAchievements()
	err := json.NewEncoder(responseWriter).Encode(achievementList)
	if err != nil {
		achievementHandler.logger.Println("[ERROR] serializing achievement", err)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetAchievementByID returns a single achievement from the database
func (achievementHandler *AchievementsHandler) GetAchievementByID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getAchievementID(request)

	achievementHandler.logger.Println("[DEBUG] getting id", id)

	achievement, err := data.GetAchievementByID(id)
	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(achievement)
		if err != nil {
			achievementHandler.logger.Println("[ERROR] serializing achievement", err)
		}
	case data.ErrorAchievementNotFound:
		achievementHandler.logger.Println("[ERROR] fetching achievement", err)
		http.Error(responseWriter, "Achievement not found", http.StatusBadRequest)
		return
	default:
		achievementHandler.logger.Println("[ERROR] fetching achievement", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}

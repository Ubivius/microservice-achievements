package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/data"
)

// GetAchievements returns the full list of achievements
func (achievementHandler *AchievementsHandler) GetAchievements(responseWriter http.ResponseWriter, request *http.Request) {
	achievementHandler.logger.Println("Handle GET achievements")
	achievementList := data.GetAchievements()
	err := data.ToJSON(achievementList, responseWriter)
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
		err = data.ToJSON(achievement, responseWriter)
		if err != nil {
			achievementHandler.logger.Println("[ERROR] serializing product", err)
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
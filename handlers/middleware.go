package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-achievements/data"
)

// MiddlewareAchievementValidation is used to validate incoming product JSONS
func (achievementHandler *AchievementsHandler) MiddlewareAchievementValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		achievement := &data.Achievement{}

		err := data.FromJSON(achievement, request.Body)
		if err != nil {
			achievementHandler.logger.Println("[ERROR] deserializing achievement", err)
			http.Error(responseWriter, "Error reading achievement", http.StatusBadRequest)
			return
		}

		// validate the achievement
		err = achievement.ValidateAchievement()
		if err != nil {
			achievementHandler.logger.Println("[ERROR] validating achievement", err)
			http.Error(responseWriter, fmt.Sprintf("Error validating achievement: %s", err), http.StatusBadRequest)
			return
		}

		// Add the achievement to the context
		context := context.WithValue(request.Context(), KeyAchievement{}, achievement)
		newRequest := request.WithContext(context)

		// Call the next handler, which can be another middleware or the final handler
		next.ServeHTTP(responseWriter, newRequest)
	})
}
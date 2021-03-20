package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/data"
)

// MiddlewareAchievementValidation is used to validate incoming achievement JSONS
func (achievementHandler *AchievementsHandler) MiddlewareAchievementValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		achievement := &data.Achievement{}

		err := json.NewDecoder(request.Body).Decode(achievement)
		if err != nil {
			log.Error(err, "Error deserializing achievement")
			http.Error(responseWriter, "Error reading achievement", http.StatusBadRequest)
			return
		}

		// validate the achievement
		err = achievement.ValidateAchievement()
		if err != nil {
			log.Error(err, "Error validating achievement")
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

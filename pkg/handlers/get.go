package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/data"
	"go.opentelemetry.io/otel"
)

// GetAchievements returns the full list of achievements
func (achievementHandler *AchievementsHandler) GetAchievements(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("template").Start(request.Context(), "getAchievements")
	defer span.End()
	log.Info("GetAchievements request")
	achievementList := achievementHandler.db.GetAchievements()
	err := json.NewEncoder(responseWriter).Encode(achievementList)
	if err != nil {
		log.Error(err, "Error serializing achievement")
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetAchievementByID returns a single achievement from the database
func (achievementHandler *AchievementsHandler) GetAchievementByID(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("template").Start(request.Context(), "getAchievementById")
	defer span.End()
	id := getAchievementID(request)

	log.Info("GetAchievementByID request", "id", id)

	achievement, err := achievementHandler.db.GetAchievementByID(id)

	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(achievement)
		if err != nil {
			log.Error(err, "Error serializing achievement")
		}
		return
	case data.ErrorAchievementNotFound:
		log.Error(err, "Achievement not found")
		http.Error(responseWriter, "Achievement not found", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error getting achievement")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}

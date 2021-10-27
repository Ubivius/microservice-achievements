package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/data"
	"go.opentelemetry.io/otel"
)

// Delete a achievement with specified id from the database
func (achievementHandler *AchievementsHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("template").Start(request.Context(), "deleteAchievementById")
	defer span.End()
	id := getAchievementID(request)
	log.Info("Delete achievement by ID request", "id", id)

	err := achievementHandler.db.DeleteAchievement(request.Context(), id)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorAchievementNotFound:
		log.Error(err, "Error deleting achievement, id does not exist")
		http.Error(responseWriter, "Achievement not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error deleting achievement")
		http.Error(responseWriter, "Error deleting achievement", http.StatusInternalServerError)
		return
	}
}

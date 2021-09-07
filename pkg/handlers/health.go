package handlers

import (
	"net/http"
)

// LivenessCheck determine when the application needs to be restarted
func (achievementHandler *AchievementsHandler) LivenessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
}

//ReadinessCheck verifies that the application is ready to accept requests
func (achievementHandler *AchievementsHandler) ReadinessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	err := achievementHandler.db.PingDB()
	
	if err != nil {
		log.Error(err, "DB unavailable")
		http.Error(responseWriter, "DB unavailable", http.StatusServiceUnavailable)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}

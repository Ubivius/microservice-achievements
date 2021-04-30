package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/database"
	"github.com/gorilla/mux"
)

// KeyAchievement is a key used for the Achievement object inside context
type KeyAchievement struct{}

// AchievementsHandler used for getting and updating achievements
type AchievementsHandler struct {
	db database.AchievementDB
}

// NewAchievementsHandler returns a pointer to a AchievementsHandler with the logger passed as a parameter
func NewAchievementsHandler(db database.AchievementDB) *AchievementsHandler {
	return &AchievementsHandler{db}
}

// getAchievementID extracts the achievement ID from the URL
// The verification of this variable is handled by gorilla/mux
func getAchievementID(request *http.Request) string {
	vars := mux.Vars(request)
	id := vars["id"]
	
	return id
}

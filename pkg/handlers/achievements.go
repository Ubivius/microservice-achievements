package handlers

import (
	"log"
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/database"
	"github.com/gorilla/mux"
)

// KeyAchievement is a key used for the Achievement object inside context
type KeyAchievement struct{}

// AchievementsHandler used for getting and updating achievements
type AchievementsHandler struct {
	logger *log.Logger
	db     database.AchievementDB
}

// NewAchievementsHandler returns a pointer to a AchievementsHandler with the logger passed as a parameter
func NewAchievementsHandler(logger *log.Logger, db database.AchievementDB) *AchievementsHandler {
	return &AchievementsHandler{logger, db}
}

// getAchievementID extracts the achievement ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getAchievementID(request *http.Request) string {
	vars := mux.Vars(request)
	id := vars["id"]
	
	return id
}

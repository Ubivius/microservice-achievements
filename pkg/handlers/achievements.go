package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyAchievement is a key used for the Achievement object inside context
type KeyAchievement struct{}

// AchievementsHandler used for getting and updating achievements
type AchievementsHandler struct {
	logger *log.Logger
}

// NewAchievementsHandler returns a pointer to a AchievementsHandler with the logger passed as a parameter
func NewAchievementsHandler(logger *log.Logger) *AchievementsHandler {
	return &AchievementsHandler{logger}
}

// getAchievementID extracts the achievement ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getAchievementID(request *http.Request) int {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
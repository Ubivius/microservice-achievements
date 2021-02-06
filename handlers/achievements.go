package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyAchievement is a key used for the Achievement object inside context
type KeyAchievement struct{}

// Achievement handler used for getting and updating achievements
type AchievementsHandler struct {
	logger *log.Logger
}

func NewAchievementsHandler(logger *log.Logger) *AchievementsHandler {
	return &AchievementsHandler{logger}
}

// getAchievementId extracts the achievement ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getAchievementId(request *http.Request) int {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
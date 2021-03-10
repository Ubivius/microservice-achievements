package router

import (
	"log"
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/handlers"
	"github.com/gorilla/mux"
)

// New : Mux route handling with gorilla/mux
func New(achievementHandler *handlers.AchievementsHandler, logger *log.Logger) *mux.Router {
	router := mux.NewRouter()

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/achievements", achievementHandler.GetAchievements)
	getRouter.HandleFunc("/achievements/{id:[0-9]+}", achievementHandler.GetAchievementByID)

	// Put router
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/achievements", achievementHandler.UpdateAchievements)
	putRouter.Use(achievementHandler.MiddlewareAchievementValidation)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/achievements", achievementHandler.AddAchievement)
	postRouter.Use(achievementHandler.MiddlewareAchievementValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/achievements/{id:[0-9]+}", achievementHandler.Delete)

	return router
}

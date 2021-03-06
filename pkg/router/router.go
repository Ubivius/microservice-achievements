package router

import (
	"net/http"

	"github.com/Ubivius/microservice-achievements/pkg/handlers"
	"github.com/gorilla/mux"
)

// New : Mux route handling with gorilla/mux
func New(achievementHandler *handlers.AchievementsHandler) *mux.Router {
	log.Info("Starting router")
	router := mux.NewRouter()

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/achievements", achievementHandler.GetAchievements)
	getRouter.HandleFunc("/achievements/{id:[0-9a-z-]+}", achievementHandler.GetAchievementByID)

	//Health Check
	getRouter.HandleFunc("/health/live", achievementHandler.LivenessCheck)
	getRouter.HandleFunc("/health/ready", achievementHandler.ReadinessCheck)

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
	deleteRouter.HandleFunc("/achievements/{id:[0-9a-z-]+}", achievementHandler.Delete)

	return router
}

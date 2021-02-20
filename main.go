package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ubivius/microservice-achievements/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Logger
	logger := log.New(os.Stdout, "Achievements", log.LstdFlags)

	// Creating handlers
	achievementHandler := handlers.NewAchievementsHandler(logger)

	// Mux route handling with gorilla/mux
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

	// Server setup
	server := &http.Server{
		Addr:        ":9090",
		Handler:     router,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Starting server on port ", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			logger.Println("Error starting server : ", err)
			logger.Fatal(err)
		}
	}()

	// Handle shutdown signals from operating system
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	receivedSignal := <-signalChannel

	logger.Println("Received terminate, beginning graceful shutdown", receivedSignal)

	// Server shutdown
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_ = server.Shutdown(timeoutContext)
}

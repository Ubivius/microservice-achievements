package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-achievements/pkg/data"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func TestValidationMiddlewareWithValidBody(t *testing.T) {
	// Creating request body
	body := &data.Achievement{
		Name:        "addName",
		Description: "addDescription",
		Condition:   "addCondition",
		SpriteID:    uuid.NewString(),
	}
	bodyBytes, _ := json.Marshal(body)

	request := httptest.NewRequest(http.MethodPost, "/achievements", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	achievementHandler := NewAchievementsHandler(newAchievementDB())

	// Create a router for middleware because function attachment is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/achievements", achievementHandler.AddAchievement)
	router.Use(achievementHandler.MiddlewareAchievementValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestValidationMiddlewareWithNoName(t *testing.T) {
	// Creating request body
	body := &data.Achievement{
		Description: "addDescription",
		Condition:   "addCondition",
		SpriteID:    uuid.NewString(),
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Error("Body passing to test is not a valid json struct : ", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/achievements", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	achievementHandler := NewAchievementsHandler(newAchievementDB())

	// Create a router for middleware because linking is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/achievements", achievementHandler.AddAchievement)
	router.Use(achievementHandler.MiddlewareAchievementValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Field validation for 'Name' failed on the 'required' tag") {
		t.Error("Expected error on field validation for Name but got : ", response.Body.String())
	}
}

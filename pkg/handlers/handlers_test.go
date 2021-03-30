package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-achievements/pkg/data"
	"github.com/Ubivius/microservice-achievements/pkg/database"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func newAchievementDB() database.AchievementDB {
	return database.NewMockAchievements()
}

func TestGetAchievements(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/achievements", nil)
	response := httptest.NewRecorder()

	achievementHandler := NewAchievementsHandler(newAchievementDB())
	achievementHandler.GetAchievements(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status code 200 but got : %d", response.Code)
	}

	if !strings.Contains(response.Body.String(), "a2181017-5c53-422b-b6bc-036b27c04fc8") || !strings.Contains(response.Body.String(), "e2382ea2-b5fa-4506-aa9d-d338aa52af44") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetExistingAchievementByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/achievements/1", nil)
	response := httptest.NewRecorder()

	achievementHandler := NewAchievementsHandler(newAchievementDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	achievementHandler.GetAchievementByID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "a2181017-5c53-422b-b6bc-036b27c04fc8") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingAchievementByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/achievements/4", nil)
	response := httptest.NewRecorder()

	achievementHandler := NewAchievementsHandler(newAchievementDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": uuid.NewString(),
	}
	request = mux.SetURLVars(request, vars)

	achievementHandler.GetAchievementByID(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Achievement not found") {
		t.Error("Expected response : Achievement not found")
	}
}

func TestDeleteNonExistantAchievement(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/achievements/4", nil)
	response := httptest.NewRecorder()

	achievementHandler := NewAchievementsHandler(newAchievementDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": uuid.NewString(),
	}
	request = mux.SetURLVars(request, vars)

	achievementHandler.Delete(response, request)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Achievement not found") {
		t.Error("Expected response : Achievement not found")
	}
}

func TestAddAchievement(t *testing.T) {
	// Creating request body
	body := &data.Achievement{
		Name:        "addName",
		Description: "addDescription",
		Condition:   "addCondition",
		SpriteID:    1,
	}

	request := httptest.NewRequest(http.MethodPost, "/achievements", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyAchievement{}, body)
	request = request.WithContext(ctx)

	achievementHandler := NewAchievementsHandler(newAchievementDB())
	achievementHandler.AddAchievement(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestUpdateNonExistantAchievement(t *testing.T) {
	// Creating request body
	body := &data.Achievement{
		ID:          uuid.NewString(),
		Name:        "addName",
		Description: "addDescription",
		Condition:   "addCondition",
		SpriteID:    1,
	}

	request := httptest.NewRequest(http.MethodPut, "/achievements", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyAchievement{}, body)
	request = request.WithContext(ctx)

	achievementHandler := NewAchievementsHandler(newAchievementDB())
	achievementHandler.UpdateAchievements(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, response.Code)
	}
}

func TestUpdateAchievement(t *testing.T) {
	// Creating request body
	body := &data.Achievement{
		ID:          "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Name:        "addName",
		Description: "addDescription",
		Condition:   "addCondition",
		SpriteID:    1,
	}

	request := httptest.NewRequest(http.MethodPut, "/achievements", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyAchievement{}, body)
	request = request.WithContext(ctx)

	achievementHandler := NewAchievementsHandler(newAchievementDB())
	achievementHandler.UpdateAchievements(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteExistingAchievement(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/achievements/1", nil)
	response := httptest.NewRecorder()

	achievementHandler := NewAchievementsHandler(newAchievementDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	achievementHandler.Delete(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}

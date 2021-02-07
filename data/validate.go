package data

import (
	"github.com/go-playground/validator"
)

// The current setup works well with a single struct to validate
// The struct to validate should be passed as an interface in the future and the errors should be handled as individual error strings
// For further information see :
// Validator library : https://github.com/go-playground/validator
// Nic Jackson episode : https://github.com/nicholasjackson/building-microservices-youtube/blob/episode_7/product-api/data/validation.go

// ValidateAchievement a achievement with json validation
func (achievement *Achievement) ValidateAchievement() error {
	validate := validator.New()

	return validate.Struct(achievement)
}

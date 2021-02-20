package data

import (
	"github.com/go-playground/validator"
)

// ValidateAchievement a achievement with json validation
func (achievement *Achievement) ValidateAchievement() error {
	validate := validator.New()

	return validate.Struct(achievement)
}

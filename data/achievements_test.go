package data

import "testing"

func TestChecksValidation(t *testing.T) {
	achievement := &Achievement{
		Name:  "Malcolm",
		Price: 2.00,
		SKU:   "abs-abs-abscd",
	}

	err := achievement.ValidateAchievement()

	if err != nil {
		t.Fatal(err)
	}
}
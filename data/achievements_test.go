package data

import "testing"

func TestChecksValidation(t *testing.T) {
	achievement := &Achievement{
		Name: "100 kills",
	}

	err := achievement.ValidateAchievement()

	if err != nil {
		t.Fatal(err)
	}
}

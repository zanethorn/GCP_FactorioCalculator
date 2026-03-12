package validation

import (
	"errors"
	. "shared/models"
)

func ValidateRecipe(r Recipe) error {
	if r.Name == "" {
		return errors.New("name required")
	}
	if len(r.Ingredients) == 0 {
		return errors.New("at least one ingredient required")
	}
	return nil
}

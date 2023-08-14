package validator_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/max-weis/go-asyncapi-validator/pkg/validator"
)

func TestValidateJSONAgainstSchema(t *testing.T) {
	validSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
			"age": map[string]interface{}{
				"type": "integer",
			},
		},
		"required": []string{"name"},
	}

	validJSON := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	invalidJSON := map[string]interface{}{
		"age": 30,
	}

	t.Run("valid JSON", func(t *testing.T) {
		if err := validator.ValidateJSONAgainstSchema(validJSON, validSchema); err != nil {
			t.Errorf("Expected no error but got %s", err)
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		err := validator.ValidateJSONAgainstSchema(invalidJSON, validSchema)
		if err == nil {
			t.Error("Expected an error for invalid JSON but got nil")
		} else if diff := cmp.Diff("json is not valid", err.Error()); diff != "" {
			t.Errorf("Mismatch (-expected +got):\n%s", diff)
		}
	})
}

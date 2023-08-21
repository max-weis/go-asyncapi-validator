package validator_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/max-weis/go-asyncapi-validator/pkg/validator"
	"testing"
)

func TestExtractSchemaWithJSONPath(t *testing.T) {
	mockSpec := map[string]interface{}{
		"info": map[string]interface{}{
			"title":   "Test API",
			"version": "1.0",
		},
		"channels": map[string]interface{}{
			"userUpdates": map[string]interface{}{
				"publish": map[string]interface{}{
					"message": map[string]interface{}{
						"payload": "sample payload",
					},
				},
			},
		},
		"components": map[string]interface{}{
			"schemas": map[string]interface{}{
				"testSchema": map[string]interface{}{
					"type": "object",
				},
			},
		},
	}

	t.Run("Extract title", func(t *testing.T) {
		query := "$.info.title"
		expectedResult := "Test API"
		result1, err1 := validator.ExtractSchemaWithJSONPath(mockSpec, query)
		if err1 != nil {
			t.Errorf("Expected no error but got %s", err1)
		}
		if diff := cmp.Diff(expectedResult, result1); diff != "" {
			t.Errorf("Mismatch (-expected +got):\n%s", diff)
		}
	})

	t.Run("Extract payload", func(t *testing.T) {
		query := "$.channels.userUpdates.publish.message.payload"
		expectedResult := "sample payload"
		result2, err2 := validator.ExtractSchemaWithJSONPath(mockSpec, query)
		if err2 != nil {
			t.Errorf("Expected no error but got %s", err2)
		}
		if diff := cmp.Diff(expectedResult, result2); diff != "" {
			t.Errorf("Mismatch (-expected +got):\n%s", diff)
		}
	})

	t.Run("Invalid path", func(t *testing.T) {
		query := "$.invalid.path"
		expectedError := "key error: invalid not found in object"
		_, err3 := validator.ExtractSchemaWithJSONPath(mockSpec, query)
		if err3 == nil || err3.Error() != expectedError {
			t.Errorf("Expected error '%s' but got '%v'", expectedError, err3)
		}
	})

}

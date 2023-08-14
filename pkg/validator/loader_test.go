package validator_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/max-weis/go-asyncapi-validator/pkg/validator"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadAsyncAPISpec(t *testing.T) {
	// Setup a temporary file for testing
	tempFile, err := ioutil.TempFile("", "asyncapi_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(tempFile.Name())

	mockSpec := `{
		"info": {
			"title": "Test API",
			"version": "1.0"
		}
	}`
	if _, err := tempFile.Write([]byte(mockSpec)); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}

	t.Run("load valid AsyncAPI spec", func(t *testing.T) {
		spec, err := validator.LoadAsyncAPISpec(tempFile.Name())
		if err != nil {
			t.Fatalf("Expected no error but got %s", err)
		}

		expectedSpec := map[string]interface{}{
			"info": map[string]interface{}{
				"title":   "Test API",
				"version": "1.0",
			},
		}
		if diff := cmp.Diff(expectedSpec, spec); diff != "" {
			t.Errorf("Mismatch (-expected +got):\n%s", diff)
		}
	})

	t.Run("load invalid path", func(t *testing.T) {
		_, err := validator.LoadAsyncAPISpec("invalid_path.json")
		if err == nil {
			t.Errorf("Expected error for invalid path but got nil")
		}
	})
}
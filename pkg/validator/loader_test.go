package validator

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLoadAsyncAPISpec(t *testing.T) {
	t.Run("load valid AsyncAPI spec", func(t *testing.T) {
		mockSpec := `{
			"info": {
				"title": "Test API",
				"version": "1.0"
			}
		}`

		spec, err := loadAsyncAPISpec(mockSpec)
		if err != nil {
			t.Fatalf("Expected no error but got %s", err)
		}

		expectedSpec := AsyncAPI{
			"info": map[string]interface{}{
				"title":   "Test API",
				"version": "1.0",
			},
		}
		if diff := cmp.Diff(expectedSpec, spec); diff != "" {
			t.Errorf("Mismatch (-expected +got):\n%s", diff)
		}
	})

	t.Run("load invalid AsyncAPI spec", func(t *testing.T) {
		invalidSpec := `{ "info": "Test API" `

		_, err := loadAsyncAPISpec(invalidSpec)
		if err == nil {
			t.Errorf("Expected error for invalid JSON but got nil")
		}
	})

	t.Run("Extract and dereference $ref", func(t *testing.T) {
		mockSpec := AsyncAPI{
			"info": map[string]interface{}{
				"title":   "Test API",
				"version": "1.0",
			},
			"channels": map[string]interface{}{
				"userUpdates": map[string]interface{}{
					"publish": map[string]interface{}{
						"message": map[string]interface{}{
							"payload": map[string]interface{}{
								"$ref": "#/components/schemas/testSchema",
							},
						},
					},
				},
			},
			"components": map[string]interface{}{
				"schemas": map[string]interface{}{
					"testSchema": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"sampleKey": map[string]interface{}{
								"type": "string",
							},
						},
					},
				},
			},
		}

		query := "$.channels.userUpdates.publish.message.payload"
		// Assuming that the payload field contains a $ref to components.schemas.testSchema
		expectedResult := map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"sampleKey": map[string]interface{}{
					"type": "string",
				},
			},
		}
		result4, err4 := extractSchemaWithJSONPath(mockSpec, query)
		if err4 != nil {
			t.Errorf("Expected no error but got %s", err4)
		}
		if diff := cmp.Diff(expectedResult, result4); diff != "" {
			t.Errorf("Mismatch (-expected +got):\n%s", diff)
		}
	})

}

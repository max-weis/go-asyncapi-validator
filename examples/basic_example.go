package main

import (
	"fmt"
	"github.com/max-weis/go-asyncapi-validator/pkg/validator"
)

func main() {
	// Load AsyncAPI Spec
	spec, err := validator.LoadAsyncAPISpec("./spec.json")
	if err != nil {
		fmt.Printf("Failed to load AsyncAPI spec: %s", err)
		return
	}

	// Extract Schema using JSON Path
	schema, err := validator.ExtractSchemaWithJSONPath(spec, "$.channels.personUpdates.subscribe.message.payload")
	if err != nil {
		fmt.Printf("Failed to extract schema: %s", err)
		return
	}

	// Sample JSON data you want to validate
	jsonData := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Joe",
		Age:  38,
	}

	// Validate
	err = validator.ValidateJSONAgainstSchema(jsonData, schema)
	if err != nil {
		fmt.Printf("Validation failed: %s", err)
		return
	}

	fmt.Println("Validation succeeded!")
}

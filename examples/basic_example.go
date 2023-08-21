package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/max-weis/go-asyncapi-validator/pkg/validator"
	"os"
)

func main() {
	// Load AsyncAPI Spec
	spec, err := validator.LoadAsyncAPISpecFromFile("./spec.json")
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
	jsonData, err := loadFile("./example.json")
	if err != nil {
		fmt.Printf("Failed to example json: %s", err)
		return
	}

	// Validate
	err = validator.ValidateJSONAgainstSchema(jsonData, schema)
	if err != nil {
		fmt.Printf("Validation failed: %s", err)
		var e validator.ValidationError
		ok := errors.As(err, &e)
		if !ok {
			return
		}

		fmt.Println(e.PrettyPrint())

		return
	}

	fmt.Println("Validation succeeded!")
}

func loadFile(path string) (interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var file interface{}
	err = json.Unmarshal(data, &file)
	return file, err
}

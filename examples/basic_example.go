package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/max-weis/go-asyncapi-validator/pkg/validator"
	"os"
)

func main() {
	spec, err := validator.LoadAsyncAPISpecFromFile("./spec.json")
	if err != nil {
		fmt.Printf("Failed to load AsyncAPI spec: %s", err)
		return
	}

	// or use json path
	// query := "$.channels.personUpdates.subscribe.message.payload"
	query := validator.NewBuilder().Channels("personUpdates").Subscribe().Payload()
	schema, err := validator.ExtractSchemaWithJSONPath(spec, query)
	if err != nil {
		fmt.Printf("Failed to extract schema: %s", err)
		return
	}

	jsonData, err := loadFile("./example.json")
	if err != nil {
		fmt.Printf("Failed to example json: %s", err)
		return
	}

	if err = validator.ValidateJSONAgainstSchema(jsonData, schema); err != nil {
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

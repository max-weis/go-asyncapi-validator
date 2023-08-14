package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/oliveagle/jsonpath"
	"github.com/xeipuuv/gojsonschema"
)

func loadAsyncAPISpec(path string) (interface{}, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var spec interface{}
	err = json.Unmarshal(data, &spec)
	return spec, err
}

func extractSchemaWithJSONPath(spec interface{}, query string) (interface{}, error) {
	res, err := jsonpath.JsonPathLookup(spec, query)
	return res, err
}

func validateJSONAgainstSchema(jsonData interface{}, schema interface{}) error {
	schemaLoader := gojsonschema.NewGoLoader(schema)
	documentLoader := gojsonschema.NewGoLoader(jsonData)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		for _, err := range result.Errors() {
			fmt.Printf("- %s\n", err)
		}
		return fmt.Errorf("json is not valid")
	}
	return nil
}

func main() {
	// Load AsyncAPI Spec
	spec, err := loadAsyncAPISpec("./spec.json")
	if err != nil {
		log.Fatalf("Failed to load AsyncAPI spec: %s", err)
	}

	// Extract Schema using JSON Path
	schema, err := extractSchemaWithJSONPath(spec, "$.channels.personUpdates.subscribe.message.payload")
	if err != nil {
		log.Fatalf("Failed to extract schema: %s", err)
	}

	// Load JSON data you want to validate
	jsonData := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Max",
		Age:  24,
	}

	// Validate
	err = validateJSONAgainstSchema(jsonData, schema)
	if err != nil {
		log.Fatalf("Validation failed: %s", err)
	}

	fmt.Println("Validation succeeded!")
}

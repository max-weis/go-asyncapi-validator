package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/max-weis/go-asyncapi-validator/pkg/validator"
	"log"
	"os"
)

func main() {
	// Define CLI flags
	pathToSpec := flag.String("spec", "", "Path to the AsyncAPI spec file.")
	pathToJson := flag.String("json", "", "Path to the JSON file to be validated.")
	jsonPath := flag.String("jsonpath", "", "The location of the schema inside the spec to which the json will be validated agains")

	// Parse CLI flags
	flag.Parse()

	if *pathToSpec == "" || *pathToJson == "" || *jsonPath == "" {
		fmt.Println("spec, json and jsonpath parameters are required.")
		flag.PrintDefaults()
		return
	}

	// Load AsyncAPI Spec
	spec, err := validator.LoadAsyncAPISpecFromFile(*pathToSpec)
	if err != nil {
		log.Fatalf("Failed to load AsyncAPI spec: %s", err)
	}

	// Extract Schema using JSON Path
	schema, err := validator.ExtractSchemaWithJSONPath(spec, *jsonPath)
	if err != nil {
		log.Fatalf("Failed to extract schema: %s", err)
	}

	// Load JSON data from provided path
	jsonData, err := loadFile(*pathToJson)
	if err != nil {
		log.Fatalf("Failed to load JSON data: %s", err)
	}

	// Validate
	if err = validator.ValidateJSONAgainstSchema(jsonData, schema); err != nil {
		log.Fatalf("Validation failed: %s", err)
	}

	fmt.Println("the provided JSON is valid")
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

package validator

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
)

type AsyncAPI map[string]any

// LoadAsyncAPISpecFromFile loads and parses an AsyncAPI specification from a given file path.
//
// Parameters:
//   - path: The path to the file containing the AsyncAPI specification, in YAML/JSON format.
//
// Returns:
//   - The parsed AsyncAPI specification.
//   - An error if there are issues reading the file or parsing the contained YAML/JSON.
//
// Important considerations:
//  1. The function expects the file at the provided path to contain a valid YAML/JSON representation of an AsyncAPI spec.
//  2. The returned interface{} is typically a map[string]interface{} for JSON objects or a slice for JSON arrays.
//     Type assertion might be necessary based on the structure of your AsyncAPI spec.
//  3. Errors might arise from file access issues (e.g., file not found, permission issues) or YAML/JSON parsing issues
//     (e.g., malformed YAML/JSON, unexpected data types).
//  4. Ensure that the provided path is either an absolute path or relative to the current working directory of the executable.
func LoadAsyncAPISpecFromFile(path string) (AsyncAPI, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return unmarshalIntoMap(data)
}

// LoadAsyncAPISpec loads and parses an AsyncAPI specification from the given spec.
//
// Parameters:
//   - spec: the AsyncAPI specification, in JSON format.
//
// Returns:
//   - The parsed AsyncAPI specification.
//   - An error if there are issues reading the file or parsing the contained YAML/JSON.
//
// Important considerations:
//  1. The function expects the file at the provided path to contain a valid YAML/JSON representation of an AsyncAPI spec.
//  2. The returned interface{} is typically a map[string]interface{} for JSON objects or a slice for JSON arrays.
//     Type assertion might be necessary based on the structure of your AsyncAPI spec.
//  3. Errors might arise from file access issues (e.g., file not found, permission issues) or YAML/JSON parsing issues
//     (e.g., malformed YAML/JSON, unexpected data types).
//  4. Ensure that the provided path is either an absolute path or relative to the current working directory of the executable.
func LoadAsyncAPISpec(spec string) (AsyncAPI, error) {
	return unmarshalIntoMap([]byte(spec))
}

func unmarshalIntoMap(spec []byte) (AsyncAPI, error) {
	var doc map[string]interface{}
	if err := json.Unmarshal(spec, &doc); err != nil {
		if err := yaml.Unmarshal(spec, &doc); err != nil {
			return nil, err
		}
	}

	return doc, nil
}

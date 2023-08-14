package validator

import (
	"encoding/json"
	"os"
)

// LoadAsyncAPISpec loads and parses an AsyncAPI specification from a given file path.
//
// Parameters:
//   - path: The path to the file containing the AsyncAPI specification, typically in JSON format.
//
// Returns:
//   - A generic interface{} representing the parsed AsyncAPI specification.
//   - An error if there are issues reading the file or parsing the contained JSON.
//
// Important considerations:
//  1. The function expects the file at the provided path to contain a valid JSON representation of an AsyncAPI spec.
//  2. The returned interface{} is typically a map[string]interface{} for JSON objects or a slice for JSON arrays.
//     Type assertion might be necessary based on the structure of your AsyncAPI spec.
//  3. Errors might arise from file access issues (e.g., file not found, permission issues) or JSON parsing issues
//     (e.g., malformed JSON, unexpected data types).
//  4. Ensure that the provided path is either an absolute path or relative to the current working directory of the executable.
func LoadAsyncAPISpec(path string) (interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var spec interface{}
	err = json.Unmarshal(data, &spec)
	return spec, err
}

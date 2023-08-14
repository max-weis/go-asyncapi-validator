package validator

import "github.com/oliveagle/jsonpath"

// ExtractSchemaWithJSONPath retrieves a part of a given spec based on a JSON Path query.
//
// Parameters:
//   - spec: The data structure (usually from a parsed JSON document) from which a part needs to be extracted.
//   - query: The JSON Path query string used to specify the part of the spec to extract.
//
// Returns:
//   - The extracted data based on the provided JSON Path query.
//   - An error if the JSON Path query is invalid, if there's no match, or if there are other issues during extraction.
//
// Important considerations:
//  1. The function uses the 'github.com/oliveagle/jsonpath' library to perform the extraction.
//     Familiarity with its syntax and constraints is recommended.
//  2. The given spec should preferably be a map or slice. Providing simple data types might result in unexpected behaviors.
//  3. If the query does not match any part of the spec, an error indicating "unknown key" will be returned.
func ExtractSchemaWithJSONPath(spec interface{}, query string) (interface{}, error) {
	return jsonpath.JsonPathLookup(spec, query)
}

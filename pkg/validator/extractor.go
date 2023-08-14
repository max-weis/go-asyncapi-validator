package validator

import "github.com/oliveagle/jsonpath"

func ExtractSchemaWithJSONPath(spec interface{}, query string) (interface{}, error) {
	return jsonpath.JsonPathLookup(spec, query)
}

package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/oliveagle/jsonpath"
)

func ExtractSchemaWithJSONPath(spec map[string]interface{}, query string) (interface{}, error) {
	value, err := jsonpath.JsonPathLookup(spec, query)
	if err != nil {
		return nil, err
	}

	schema, ok := value.(map[string]interface{})
	if !ok {
		return value, nil
	}

	ref, ok := schema["$ref"].(string)
	if !ok || !strings.HasPrefix(ref, "#/components/schemas/") {
		return value, nil
	}

	schemaName := strings.TrimPrefix(ref, "#/components/schemas/")

	components, ok := spec["components"].(map[string]interface{})
	if !ok {
		return nil, errors.New("components not found in the spec")
	}

	schemas, ok := components["schemas"].(map[string]interface{})
	if !ok {
		return nil, errors.New("components.schemas not found in the spec")
	}

	derefSchema, ok := schemas[schemaName].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("schema '%s' not found in components.schemas", schemaName)
	}

	if _, hasRef := derefSchema["$ref"]; hasRef {
		return ExtractSchemaWithJSONPath(spec, fmt.Sprintf("$.components.schemas.%s", schemaName))
	}

	return derefSchema, nil
}

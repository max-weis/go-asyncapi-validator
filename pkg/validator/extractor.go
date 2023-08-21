package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/oliveagle/jsonpath"
)

func ExtractSchemaWithJSONPath(spec AsyncAPI, query string) (interface{}, error) {
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

type Builder struct {
	pathSegments []string
}

func NewBuilder() *Builder {
	return &Builder{
		pathSegments: []string{"$"},
	}
}

func (b *Builder) Channels(key string) *Builder {
	b.pathSegments = append(b.pathSegments, "channels", key)
	return b
}

func (b *Builder) Subscribe() *Builder {
	b.pathSegments = append(b.pathSegments, "subscribe")
	return b
}

func (b *Builder) Publish() *Builder {
	b.pathSegments = append(b.pathSegments, "publish")
	return b
}

func (b *Builder) Payload() string {
	b.pathSegments = append(b.pathSegments, "message")
	b.pathSegments = append(b.pathSegments, "payload")
	return strings.Join(b.pathSegments, ".")
}

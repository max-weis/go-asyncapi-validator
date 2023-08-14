package validator

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateJSONAgainstSchema(jsonData interface{}, schema interface{}) error {
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

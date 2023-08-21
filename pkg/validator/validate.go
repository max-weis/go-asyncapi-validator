package validator

import (
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

type ValidationError struct {
	error
	Errs []error
}

func newValidationError(err string) ValidationError {
	return ValidationError{error: errors.New(err), Errs: make([]error, 0)}
}

func (v *ValidationError) AddErr(err string) {
	v.Errs = append(v.Errs, errors.New(err))
}

func (v *ValidationError) PrettyPrint() string {
	var message string

	for _, err := range v.Errs {
		message += fmt.Sprintf("\n- %s", err)
	}

	return message
}

// ValidateJSONAgainstSchema checks if the provided JSON data adheres to a given schema.
//
// Parameters:
//   - jsonData: The data (usually a map or slice from parsed JSON) that needs validation against the schema.
//   - schema: The schema (typically in JSON format) to which jsonData should adhere.
//
// Returns:
//   - nil if the jsonData matches the schema without any issues.
//   - An error detailing the mismatch or any other issues encountered during validation.
//
// Important considerations:
//  1. The function uses the 'github.com/xeipuuv/gojsonschema' library for JSON schema validation. It supports the
//     JSON Schema Draft 4, 6, and 7 specifications.
//  2. Both the jsonData and schema parameters should preferably be of type map[string]interface{} or appropriate
//     Go types representing JSON structures.
//  3. If the validation fails, the function prints each validation error to the standard output.
//  4. It's recommended to handle and process the returned error appropriately in your application. If validation fails,
//     the error will have the message "json is not valid".
//  5. Ensure that the provided schema is a valid JSON schema; otherwise, the function might return unexpected errors.
func ValidateJSONAgainstSchema(jsonData interface{}, schema interface{}) error {
	schemaLoader := gojsonschema.NewGoLoader(schema)
	documentLoader := gojsonschema.NewGoLoader(jsonData)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		validationErr := newValidationError("json is not valid")
		for _, err := range result.Errors() {
			validationErr.AddErr(err.String())
		}
		return validationErr
	}

	return nil
}

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

type Validator struct {
	spec string
	path string
}

func NewValidator(spec string, path string) Validator {
	return Validator{
		spec: spec,
		path: path,
	}
}

func (v *Validator) Validate(object interface{}) (bool, error) {
	spec, err := loadAsyncAPISpec(v.spec)
	if err != nil {
		return false, err
	}

	schema, err := extractSchemaWithJSONPath(spec, v.path)
	if err != nil {
		return false, err
	}

	if err = validateJSONAgainstSchema(object, schema); err != nil {
		return false, err
	}

	return true, nil
}

func validateJSONAgainstSchema(jsonData interface{}, schema interface{}) error {
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

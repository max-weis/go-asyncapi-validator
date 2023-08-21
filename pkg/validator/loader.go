package validator

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

type AsyncAPI map[string]any

func loadAsyncAPISpec(spec string) (AsyncAPI, error) {
	var doc map[string]interface{}
	if err := json.Unmarshal([]byte(spec), &doc); err != nil {
		if err := yaml.Unmarshal([]byte(spec), &doc); err != nil {
			return nil, err
		}
	}

	return doc, nil
}

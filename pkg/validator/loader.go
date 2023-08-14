package validator

import (
	"encoding/json"
	"os"
)

func LoadAsyncAPISpec(path string) (interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var spec interface{}
	err = json.Unmarshal(data, &spec)
	return spec, err
}

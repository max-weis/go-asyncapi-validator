package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/max-weis/go-asyncapi-validator/pkg/validator"

	"os"
)

func main() {
	spec, err := os.ReadFile("./spec.json")
	if err != nil {
		fmt.Printf("Failed to load AsyncAPI spec: %s", err)
		return
	}

	file, err := os.ReadFile("./example.json")
	if err != nil {
		fmt.Printf("Failed to load json: %s", err)
		return
	}

	var obj interface{}
	if err = json.Unmarshal(file, &obj); err != nil {
		fmt.Printf("Failed to parse json: %s", err)
		return
	}

	// or use json path
	// query := "$.channels.personUpdates.subscribe.message.payload"
	query := validator.NewBuilder().Channels("personUpdates").Subscribe().Payload()
	v := validator.NewValidator(string(spec), query)
	ok, err := v.Validate(obj)
	if err != nil {
		fmt.Printf("Validation failed: %s", err)
		var e validator.ValidationError
		ok := errors.As(err, &e)
		if !ok {
			return
		}

		fmt.Println(e.PrettyPrint())

		return
	}

	if ok {
		fmt.Println("Validation succeeded!")
	}
}

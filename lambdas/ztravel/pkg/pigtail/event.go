package pigtail

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type MyEvent struct {
	Name string `json:"name,omitempty" validate:"required,nonblank"`
}

func NewMyEventFromJSON(validate *validator.Validate, value string) (*MyEvent, error) {
	var event MyEvent

	err := json.Unmarshal([]byte(value), &event)
	if err != nil {
		return nil, err
	}

	err = validate.Struct(&event)
	if err != nil {
		return nil, err
	}

	return &event, err
}

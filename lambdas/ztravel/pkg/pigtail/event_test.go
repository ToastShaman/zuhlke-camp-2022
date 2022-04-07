package pigtail

import (
	"testing"
	"ztravel/pkg/validation"

	"github.com/stretchr/testify/assert"
)

func TestNewMyEventFromJSON(t *testing.T) {
	v := validation.New()

	t.Run("Valid", func(t *testing.T) {
		event, err := NewMyEventFromJSON(v, `{ "name": "Emma" }`)
		assert.Nil(t, err)
		assert.Equal(t, event, &MyEvent{Name: "Emma"})
	})

	t.Run("Invalid (Spaces)", func(t *testing.T) {
		event, err := NewMyEventFromJSON(v, `{ "name": " " }`)
		assert.Nil(t, event)
		assert.NotNil(t, err)
	})

	t.Run("Invalid (Emtpy)", func(t *testing.T) {
		event, err := NewMyEventFromJSON(v, `{ "name": "" }`)
		assert.Nil(t, event)
		assert.NotNil(t, err)
	})
}

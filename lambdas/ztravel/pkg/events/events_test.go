package events

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	events := NewRecordingEvents()
	event := &testEvent{ID: 1}
	events.Log(event)

	assert.IsType(t, event, events.CapturedEvents[0])
	assert.JSONEq(t, `{"level":"info","event":{"id":1},"message":"My Test Event"}`, events.Output()[0])
}

type testEvent struct {
	ID int
}

func (m testEvent) Level() zerolog.Level {
	return zerolog.InfoLevel
}

func (m testEvent) Message() string {
	return "My Test Event"
}

func (m testEvent) MarshalZerologObject(e *zerolog.Event) {
	e.Int("id", m.ID)
}

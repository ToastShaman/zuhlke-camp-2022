package events

import (
	"github.com/rs/zerolog"
)

type Event interface {
	zerolog.LogObjectMarshaler
	Level() zerolog.Level
	Message() string
}

type Events interface {
	Log(event Event)
}

type PrintingEvents struct {
	log zerolog.Logger
}

func (p *PrintingEvents) Log(event Event) {
	p.log.WithLevel(event.Level()).Object("event", event).Msg(event.Message())
}

func NewPrintingEvents(l zerolog.Logger) *PrintingEvents {
	return &PrintingEvents{log: l}
}

type RecordingEvents struct {
	CapturedEvents []interface{}
	delegate       Events
	capturedOutput *dummyWriter
}

func (r *RecordingEvents) Log(event Event) {
	r.delegate.Log(event)
	r.CapturedEvents = append(r.CapturedEvents, event)
}

func (r *RecordingEvents) Output() []string {
	return r.capturedOutput.Output
}

func NewRecordingEvents() *RecordingEvents {
	writer := &dummyWriter{
		Output: []string{},
	}

	log := zerolog.New(writer).With().Logger()

	return &RecordingEvents{
		CapturedEvents: make([]interface{}, 0),
		capturedOutput: writer,
		delegate:       NewPrintingEvents(log),
	}
}

type dummyWriter struct {
	Output []string
}

func (d *dummyWriter) Write(p []byte) (n int, err error) {
	d.Output = append(d.Output, string(p))
	return 0, nil
}

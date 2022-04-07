package main

import (
	"context"
	"ztravel/pkg/pigtail"
	"ztravel/pkg/validation"

	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(v *validator.Validate) func(context.Context, events.SQSEvent) (string, error) {
	fn := func(ctx context.Context, sqsEvent events.SQSEvent) (string, error) {
		for _, message := range sqsEvent.Records {
			event, err := pigtail.NewMyEventFromJSON(v, message.Body)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to unmarshall event")
			}

			log.Info().Str("name", event.Name).Msg("Hello")
		}
		return "Done", nil
	}

	return fn
}

func main() {
	v := validation.New()

	lambda.Start(Handler(v))
}

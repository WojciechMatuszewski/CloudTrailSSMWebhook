package main

import (
	"bytes"
	"errors"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"
)

func handler(events events.CloudWatchEvent) {
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		panic(errors.New("no env variable"))
	}

	raw, err := events.Detail.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	log.Info().RawJSON("inc", raw).Msg("event")

	_, err = http.Post(webhookURL, "application/json", bytes.NewBufferString(`{"text": "Parameter change!"}`))
	if err != nil {
		log.Err(err).Msg("while sending post request")
	}
}

func main() {
	lambda.Start(handler)
}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess := session.Must(session.NewSession())
	ssmc := ssm.New(sess)

	ssmK := os.Getenv("SSM_KEY_PATH")
	if ssmK == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "env var not found",
		}, nil
	}

	_, err := ssmc.PutParameter(&ssm.PutParameterInput{
		Name:      aws.String(ssmK),
		Overwrite: aws.Bool(true),
		Value:     aws.String(fmt.Sprintf("%v", rand.Intn(500))),
		Type:      aws.String("String"),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       http.StatusText(http.StatusOK),
	}, nil
}

func main() {
	lambda.Start(handler)
}

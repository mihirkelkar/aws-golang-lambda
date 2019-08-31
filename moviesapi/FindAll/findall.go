package main

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Movie struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var movies = []Movie{
	Movie{ID: 1, Name: "Avengers"},
	Movie{ID: 2, Name: "Ant-Man"},
	Movie{ID: 3, Name: "Thor"},
	Movie{ID: 4, Name: "Hulk"},
	Movie{ID: 5, Name: "Doctor Strange"},
}

func FindAll() (events.APIGatewayProxyResponse, error) {
	response, err := json.Marshal(movies)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(response),
	}, nil
}

func FindOne(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(req.PathParameters["id"])

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "ID is not a number",
		}, nil
	}

	response, err := json.Marshal(movies[id-1])

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "ID is not a number",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil

}

func main() {
	lambda.Start(FindAll)
}

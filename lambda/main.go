package main

import (
	"context"
	"log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/konstantinfarrell/go-example-lambda/pkg/file/handlers"
)

func Handler(ctx context.Context, event events.KinesisEvent) {
	log.Printf("Configure DB")
	fileHandler, err := file.New()
	if err != nil {
		panic(err)
	}

	log.Printf("Handle records")
	for _, record := range event.Records {
		kinesisRecord := record.Kinesis
		dataBytes := kinesisRecord.Data
		dataText := string(dataBytes)
		log.Printf("Handle record: %s", dataText)
		fileHandler.Handle(dataText)
	}
}

func main(){
	log.Printf("Begin function")
	lambda.Start(Handler)
}
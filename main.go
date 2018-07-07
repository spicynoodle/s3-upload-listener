package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Handler(ctx context.Context, s3Event events.S3Event) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	// Create DynamoDB client
	awsDynamoDB := dynamodb.New(sess)
	awsRekognition := rekognition.New(sess)

	for _, record := range s3Event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		fmt.Printf("Hello From Go")

	}

	imageToParse := rekognition.DetectTextInput{
		Image: &rekognition.Image{
			S3Object: &s3.Object{
				Bucket: "",
				Name:   "",
			},
		},
	}
	plateNumber, _ := awsRekognition.DetectText(&imageToParse)
}

func main() {

	lambda.Start(Handler)
}

package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

//Handler is used to handle Lambda event
func Handler(ctx context.Context, s3Event events.S3Event) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	// Create DynamoDB client
	//awsDynamoDB := dynamodb.New(sess)
	awsRekognition := rekognition.New(sess)

	for _, record := range s3Event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)

		imageToParse := rekognition.DetectTextInput{
			Image: &rekognition.Image{
				S3Object: &rekognition.S3Object{
					Bucket: &s3.Bucket.Name,
					Name:   &s3.Object.Key,
				},
			},
		}

		rekognitionResult, err := awsRekognition.DetectText(&imageToParse)
		if err != nil {
			fmt.Println(err)
		}

		textDetectionResult := ""

		for _, v := range rekognitionResult.TextDetections {
			textDetectionResult = textDetectionResult + v.String()
		}

		regexPattern, _ := regexp.Compile(`([A-Z]{1,2}([\s]{1}|[\S])(\d){1,5}([\s]{1}|[\S])[A-Z]{2,3})`)
		result := regexPattern.FindString(textDetectionResult)

		fmt.Println(result)

	}
}

func main() {

	lambda.Start(Handler)
}

package main

import (
	"context"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func handler(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	if event.RequestType != cfn.RequestDelete {
		return
	}

	sess := session.Must(session.NewSession())
	s3Client := s3.New(sess)

	bucket, ok := event.ResourceProperties["BucketName"].(string)
	if !ok {
		return event.PhysicalResourceID, nil, err
	}


	iter := s3manager.NewDeleteListIterator(s3Client, &s3.ListObjectsInput{Bucket: aws.String(bucket)})
	err = s3manager.NewBatchDeleteWithClient(s3Client).Delete(ctx, iter)
	if err != nil {
		return event.PhysicalResourceID, nil, err
	}

	return
}


func main() {
	lambda.Start(cfn.LambdaWrap(handler))
}

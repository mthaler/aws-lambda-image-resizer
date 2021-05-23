package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/mthaler/aws-lambda-image-resizer/helpers"
)

const tmp = "/tmp/"

var sess = session.Must(session.NewSession())
var uploader = s3manager.NewUploader(sess)
var downloader = s3manager.NewDownloader(sess)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.S3Event) (string, error) {
	fmt.Printf("%v", req)
	for _, r := range req.Records {
		if key := r.S3.Object.Key; helpers.IsImage(key) {
			bucket := r.S3.Bucket.Name
			resizeImage(bucket, key)
		}
	}
	return fmt.Sprintf("%d records processed", len(req.Records)), nil
}

func resizeImage(bucket, key string) {

}
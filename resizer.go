package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/disintegration/imaging"
	"github.com/mthaler/aws-lambda-image-resizer/helpers"
	"log"
	"os"
	"path/filepath"
)

const tmp = "/tmp/"

var sess *session.Session

const (
	REGION = "eu-central-1"
	DST_BUCKET = "DST_BUCKET"
)

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	}))
}

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
	loc:= tmp + bucket + "/" + key
	dir := filepath.Dir(loc)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create temporary directory %s\n", dir)
	}

	f, err := os.Create(loc)
	if err != nil {
		log.Fatalf("Failed to create file %s\n", loc)
	}

	downloader := s3manager.NewDownloader(sess)
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatalf("Failed to download file %s from bucket %s, key %s\n", loc, bucket, key)
	}
	log.Printf("Downloaded file %s, size: %d\n", loc, n)

	src, err := imaging.Open(loc)
	if err != nil {
		log.Fatalf("Could not open image %s\n", loc)
	}

	resized := imaging.Resize(src, 200, 0, imaging.Lanczos)
	log.Printf("Resized image %s", loc)

	dst := key[:len(key)-4] + "_resized" + key[len(key)-4:]
	dstLoc := "/tmp/" + dst

	err = imaging.Save(resized, dstLoc)
	if err != nil {
		log.Fatalf("Could not save image to %s\n", dstLoc)
	}

	fup, err := os.Open(dstLoc)
	if err != nil {
		log.Fatalf("Could not open file %s\n", dstLoc)
	}

	uploader := s3manager.NewUploader(sess)

	dstBucket, ok := os.LookupEnv(DST_BUCKET)
	if !ok {
		log.Fatalf("%s environment variable not set\n", DST_BUCKET)
	} else {
		fmt.Printf("Destination bucket: %s\n", dstBucket)
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(dstBucket),
		Key: aws.String(dst),
		Body: fup,
	})

	log.Printf("Uploaded file: %s", dstLoc)
}
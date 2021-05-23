# aws-lambda-image-resizer
AWS Lambda image resizer function written in Go

## Create buckets

Open the S3 management console and create two buckets, one for the source images and one for the resized images:

![Create source bucket](screenshots/create_src_bucket.jpg)

![Create destination bucket](screenshots/create_dst_bucket.jpg)

**Important**: S3 bucket names have to be globally unique!

## Create policy

Open the IAM managment console and create a new policy:

![Create policy](screenshots/create_policy.jpg)


![Create policy](screenshots/create_policy2.jpg)

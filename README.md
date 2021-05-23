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

## Create role

Create a new role. For use case, select *Lambda*. Click on next and select the policy from the previous step:

![Attach permission policies](screenshots/attach_permissions_policies.jpg)

![Create role](screenshots/create_role.jpg)

## Build the lambda function

```bash
$ go build resizer.go
$ zip resizer.zip resizer
```
This creates an archive with the resizer binary that can be uploaded as lambda function to AWS.

## Create lambda function

Open the Lambda managment console and create a new function:

![Create function](screenshots/create_function.jpg)

This takes some time. After the function is created, the following screen shold be shown:


![Image resize function](screenshots/image_resize_function.jpg)

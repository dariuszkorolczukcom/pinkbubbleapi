package aws

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadToS3(file io.Reader, fileName string) (err error) {
	bucketName := os.Getenv("BUCKET_NAME")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: credentials.NewSharedCredentials("", "default"),
	})

	s3Svc := s3.New(sess)

	uploader := s3manager.NewUploaderWithClient(s3Svc)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucketName,
		Key:    &fileName,
		Body:   file,
	})

	return
}

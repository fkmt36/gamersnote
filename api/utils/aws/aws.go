package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/ses"
)

func NewSession() *session.Session {
	return session.New(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_KEY"),
			"",
		),
	})
}

func NewSESClient(s *session.Session) *ses.SES {
	return ses.New(s)
}

func NewS3Uploader(s *session.Session) *s3manager.Uploader {
	return s3manager.NewUploader(s)
}

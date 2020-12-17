package image

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func NewService(s3Uploader *s3manager.Uploader, bucket string) Service {
	return &service{
		uploader: s3Uploader,
		bucket:   bucket,
	}
}

// Service 画像サービスインタフェース
type Service interface {
	Upload(filename string, file io.Reader) error
}

type service struct {
	uploader *s3manager.Uploader
	bucket   string
}

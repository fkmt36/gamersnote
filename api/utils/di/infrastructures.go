package di

import (
	"gamersnote.com/v1/utils/aws"
	"gamersnote.com/v1/utils/db"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/ses"
	"gorm.io/gorm"
)

var SESClient *ses.SES
var S3Uploader *s3manager.Uploader
var DBClient *gorm.DB

func initInfra() {
	session := aws.NewSession()
	SESClient = aws.NewSESClient(session)
	S3Uploader = aws.NewS3Uploader(session)
	DBClient = db.NewDBClient()
}

package di

import (
	"os"
	"time"

	"gamersnote.com/v1/utils/tmpemail"
	"gamersnote.com/v1/utils/tmpsession"

	"gamersnote.com/v1/utils/ctxuid"
	"gamersnote.com/v1/utils/email"
	"gamersnote.com/v1/utils/image"
	"gamersnote.com/v1/utils/session"
	"gamersnote.com/v1/utils/tmpuser"
)

var CtxuidService ctxuid.Service
var EmailSender email.Sender
var ImageService image.Service
var SessionService session.Service
var TmpuserService tmpuser.Service
var TmpemailService tmpemail.Service
var TmpsessionService tmpsession.Service

func initSvc() {
	CtxuidService = ctxuid.NewService()
	EmailSender = email.NewSender(SESClient, os.Getenv("EMAIL_FROM_ADDRESS"))
	ImageService = image.NewService(S3Uploader, os.Getenv("S3_BUCKET"))
	SessionService = session.NewService(7*24*time.Hour, 12*time.Hour)
	TmpuserService = tmpuser.NewService(30*time.Minute, 60*time.Minute)
	TmpemailService = tmpemail.NewService(30*time.Minute, 60*time.Minute)
	TmpsessionService = tmpsession.NewService(30*time.Minute, 60*time.Minute)
}

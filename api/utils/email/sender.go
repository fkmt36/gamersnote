package email

import (
	"github.com/aws/aws-sdk-go/service/ses"
)

func NewSender(sesClient *ses.SES, fromAddress string) Sender {
	return &sender{
		client: sesClient,
		from:   fromAddress,
	}
}

// Sender メール送信インタフェース
type Sender interface {
	SendVerificationEmail(to string, username string, code string) error
}

type sender struct {
	client *ses.SES
	from   string
}

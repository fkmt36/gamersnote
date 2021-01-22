package email

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

func (s sender) SendPasswordResetEmail(to string, code string) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data: aws.String(fmt.Sprintf(
						"<p>パスワードを忘れた方へ</p><br><p>以下の「パスワードを再設定する」からパスワードの再設定を行ってください。</p><br><a href=\"%s/reset-password/reset?code=%s\">パスワードを再設定する</a>",
						os.Getenv("BASE_URL"),
						code,
					)),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data: aws.String(fmt.Sprintf(
						"%s/reset-password/reset?code=%s",
						os.Getenv("BASE_URL"),
						code,
					)),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("【GamersNote】パスワードの再設定"),
			},
		},
		Source: aws.String("GamersNote" + "<" + s.from + ">"),
	}
	_, err := s.client.SendEmail(input)
	if err != nil {
		return err
	}
	return nil
}

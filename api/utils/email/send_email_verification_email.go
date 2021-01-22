package email

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

func (s sender) SendEmailVerificationEmail(to string, username string, uid string, code string) error {
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
						"<p>%s様</p><br><p>メールアドレスの変更はまだ終わっていません。</p><p>以下の「確認」から変更を完了してください。</p><br><a href=\"%s/verify-email?code=%s&uid=%s\">確認</a>",
						username,
						os.Getenv("BASE_URL"),
						code,
						uid,
					)),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data: aws.String(fmt.Sprintf(
						"%s/verify-code?code=%s&uid=%s",
						os.Getenv("BASE_URL"),
						code,
						uid,
					)),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("【GamersNote】メールアドレス変更確認"),
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

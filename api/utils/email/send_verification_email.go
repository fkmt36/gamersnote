package email

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

func (s sender) SendVerificationEmail(to string, username string, code string) error {
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
						"<a href=\"%s/verify-code?code=%s&username=%s\">会員登録を完了する</a>",
						os.Getenv("BASE_URL"),
						code,
						username,
					)),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data: aws.String(fmt.Sprintf(
						"%s/verify-code?code=%s&username=%s",
						os.Getenv("BASE_URL"),
						code,
						username,
					)),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("【GamersNote】会員登録を完了してください"),
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

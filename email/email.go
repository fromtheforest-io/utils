package email

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

// SendEmailInput defines a simplified structure for sending emails.
type SendEmailInput struct {
	From    string
	To      []string
	Subject string
	Body    string
}

// SendEmailFunc is a reusable function signature that sends an email.
type SendEmailFunc func(ctx context.Context, input SendEmailInput) error

type SESAPI interface {
	SendEmail(ctx context.Context, params *ses.SendEmailInput, optFns ...func(*ses.Options)) (*ses.SendEmailOutput, error)
}

// NewSESEmailSender returns a SendEmailFunc with the SES client "closed over".
// This is similar to currying in functional languages: the SES client is baked into the returned function.
// When you call sendEmail later, it already has the client pre-bound.
func NewSESEmailSender(client SESAPI) SendEmailFunc {
	return func(ctx context.Context, input SendEmailInput) error {
		_, err := client.SendEmail(ctx, &ses.SendEmailInput{
			Source: &input.From,
			Destination: &types.Destination{
				ToAddresses: input.To,
			},
			Message: &types.Message{
				Subject: &types.Content{
					Data: &input.Subject,
				},
				Body: &types.Body{
					Html: &types.Content{
						Data: &input.Body,
					},
				},
			},
		})
		return err
	}
}

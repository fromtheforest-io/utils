package email_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/fromtheforest-io/utils/email"
	"github.com/stretchr/testify/assert"
)

// MockSESClient is a mock implementation of the SESAPI interface
type MockSESClient struct {
	CalledWith *ses.SendEmailInput
	ReturnErr  error
}

func (m *MockSESClient) SendEmail(ctx context.Context, input *ses.SendEmailInput, _ ...func(*ses.Options)) (*ses.SendEmailOutput, error) {
	m.CalledWith = input
	return &ses.SendEmailOutput{}, m.ReturnErr
}

func TestNewSESEmailSender_Success(t *testing.T) {
	mockClient := &MockSESClient{}
	send := email.NewSESEmailSender(mockClient)

	err := send(context.Background(), email.SendEmailInput{
		From:    "test@example.com",
		To:      []string{"to@example.com"},
		Subject: "Hello",
		Body:    "<p>World</p>",
	})

	assert.NoError(t, err)
	assert.NotNil(t, mockClient.CalledWith)
	assert.Equal(t, "test@example.com", *mockClient.CalledWith.Source)
	assert.Equal(t, "Hello", *mockClient.CalledWith.Message.Subject.Data)
}

func TestNewSESEmailSender_Failure(t *testing.T) {
	mockClient := &MockSESClient{
		ReturnErr: errors.New("mock error"),
	}
	send := email.NewSESEmailSender(mockClient)

	err := send(context.Background(), email.SendEmailInput{
		From:    "test@example.com",
		To:      []string{"to@example.com"},
		Subject: "Fail",
		Body:    "<p>Test</p>",
	})

	assert.Error(t, err)
	assert.EqualError(t, err, "mock error")
}

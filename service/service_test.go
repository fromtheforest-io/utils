package service

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mock Client that satisfies EventBridgeClient interface ---
type MockEventBridgeClient struct {
	mock.Mock
}

func (m *MockEventBridgeClient) PutEvents(ctx context.Context, input *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*eventbridge.PutEventsOutput), args.Error(1)
}

func TestMakeEventPublisher_Success(t *testing.T) {
	mockClient := new(MockEventBridgeClient)

	event := map[string]string{"foo": "bar"}
	eventBytes, _ := json.Marshal(event)
	detailType := "TestEvent"
	source := "my.test.source"

	// Expect a single PutEvents call with correct input
	mockClient.On("PutEvents", mock.Anything, mock.MatchedBy(func(input *eventbridge.PutEventsInput) bool {
		if len(input.Entries) != 1 {
			return false
		}
		entry := input.Entries[0]
		return *entry.DetailType == detailType &&
			*entry.Detail == string(eventBytes) &&
			*entry.Source == source
	})).Return(&eventbridge.PutEventsOutput{}, nil)

	// Act: Build and call the event publisher
	publish := MakeEventPublisher(mockClient, source)
	err := publish(context.Background(), detailType, event)

	// Assert
	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

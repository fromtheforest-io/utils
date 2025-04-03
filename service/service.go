package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	eventbridgeTypes "github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
)

type ServiceContext struct {
	PublishEvent func(ctx context.Context, detailType string, detail interface{}) error
}

type EventBridgeClient interface {
	PutEvents(ctx context.Context, input *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error)
}

func NewServiceContext(cfg aws.Config, eventSource string, ebOverride ...EventBridgeClient) *ServiceContext {
	var ebClient EventBridgeClient
	if len(ebOverride) > 0 && ebOverride[0] != nil {
		ebClient = ebOverride[0]
	} else {
		ebClient = eventbridge.NewFromConfig(cfg)
	}

	return &ServiceContext{
		PublishEvent: MakeEventPublisher(ebClient, eventSource),
	}
}

func MakeEventPublisher(ebClient EventBridgeClient, source string) func(ctx context.Context, detailType string, detail interface{}) error {
	return func(ctx context.Context, detailType string, detail interface{}) error {
		detailBytes, err := json.Marshal(detail)
		if err != nil {
			return fmt.Errorf("failed to marshal event detail: %w", err)
		}

		_, err = ebClient.PutEvents(ctx, &eventbridge.PutEventsInput{
			Entries: []eventbridgeTypes.PutEventsRequestEntry{
				{
					DetailType: aws.String(detailType),
					Detail:     aws.String(string(detailBytes)),
					Source:     aws.String(source),
				},
			},
		})
		return err
	}
}

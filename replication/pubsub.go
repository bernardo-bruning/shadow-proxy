package replication

import (
	"context"
	"shadowproxy/domain"

	"cloud.google.com/go/pubsub"
)

type PubSub struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

func NewPubSub(ctx context.Context, projectID, topicID string) (*PubSub, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	topic := client.Topic(topicID)
	topicExists, err := topic.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !topicExists {
		topic, err = client.CreateTopic(ctx, topicID)
		if err != nil {
			return nil, err
		}
	}

	return &PubSub{client: client, topic: topic}, nil
}

func (p *PubSub) Emit(ctx context.Context, req *domain.Request) error {
	message, err := req.ToMessage()
	if err != nil {
		return err
	}

	result := p.topic.Publish(ctx, message)
	_, err = result.Get(ctx)
	return err
}

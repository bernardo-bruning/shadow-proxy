package consumer

import (
	"context"
	"log"
	"net/http"

	"github.com/bernardo-bruning/shadowproxy/domain"

	"cloud.google.com/go/pubsub"
)

type Consumer struct {
	url          string
	http         http.Client
	pubsub       *pubsub.Client
	subscription *pubsub.Subscription
}

func NewConsumer(ctx context.Context, projectID, topicID, subscriptionID, url string) (*Consumer, error) {
	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	subscription := pubsubClient.Subscription(subscriptionID)
	subscriptionExists, err := subscription.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !subscriptionExists {
		topic := pubsubClient.Topic(topicID)
		subscription, err = pubsubClient.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
			Topic: topic,
		})

		if err != nil {
			return nil, err
		}
	}

	httpClient := http.Client{}
	return &Consumer{
		url:          url,
		http:         httpClient,
		pubsub:       pubsubClient,
		subscription: subscription,
	}, nil
}

func (c *Consumer) Consume(ctx context.Context) {
	c.subscription.Receive(ctx, func(_ context.Context, message *pubsub.Message) {
		req := domain.FromMessage(message)
		log.Printf("message processing: %v\n", req)
		httpReq, err := req.ToHttpRequest(c.url)
		if err != nil {
			log.Printf("failed to convert to http request:%s", err.Error())
			message.Nack()
		}

		_, err = c.http.Do(httpReq)
		if err != nil {
			log.Printf("failed to perform request:%s", err.Error())
			message.Nack()
		}
	})

}

func (c *Consumer) Close() {
	c.pubsub.Close()
}

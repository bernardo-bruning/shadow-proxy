package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/bernardo-bruning/shadowproxy/consumer"
	"github.com/bernardo-bruning/shadowproxy/filters"
	"github.com/bernardo-bruning/shadowproxy/proxy"
	"github.com/bernardo-bruning/shadowproxy/replication"
)

var consumerFlag = flag.Bool("consumer", false, "consumer")

func main() {
	flag.Parse()
	ctx := context.Background()

	if *consumerFlag {
		consume(ctx)
		return
	}

	serve(ctx)
}

func consume(ctx context.Context) {
	projectID := os.Getenv("PROJECT_ID")
	topicID := os.Getenv("TOPIC_ID")
	subscriptionID := os.Getenv("SUBSCRIPTION_ID")
	url := os.Getenv("URL")

	c, err := consumer.NewConsumer(ctx, projectID, topicID, subscriptionID, url)
	if err != nil {
		log.Fatalf("failed to start consumer:%s", err.Error())
	}
	c.Consume(ctx)

}

func serve(ctx context.Context) {
	url := os.Getenv("URL")
	port := os.Getenv("PORT")
	replicaType := os.Getenv("TYPE")

	log.Printf("redirecting to URL: %s\n", url)
	log.Printf("replication with TYPE: %s\n", replicaType)

	var replica replication.Replica

	switch replicaType {
	case "PUBSUB":
		var err error
		projectID := os.Getenv("PROJECT_ID")
		topicID := os.Getenv("TOPIC_ID")

		replica, err = replication.NewPubSub(ctx, projectID, topicID)
		if err != nil {
			log.Fatal(err)
		}
	case "LOG":
		replica = replication.NewLog()
	}

	log.Printf("starting server on PORT: %s\n", port)

	proxy, err := proxy.NewProxy(replica, url, filters.NopFilter)
	if err != nil {
		log.Fatal(err)
	}

	proxy.ListenAndServe(port)
}

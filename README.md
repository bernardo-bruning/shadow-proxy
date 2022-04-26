# Shadow proxy

A simple proxy to replicate production traffic to another infrastructure.

# Configuration
Shadow proxy is composed of 2 components, server and consumer, described your configurations bellow.

## Proxy 

You can start your proxy with environments variables:

```
shadowd
```

|Enviroment Variable               | Description                                 | Example         |
|----------------------------------|---------------------------------------------|-----------------|
| URL                              | URL target to proxy                         | http://myurl    |
| PORT                             | Port to expose proxy                        | :8080           |
| TYPE                             | Type of replication, supporte only PUBSUB   | PUBSUB          |
| PROJECT_ID                       | Google cloud PROJECT_ID                     | my-project      |
| TOPIC_ID                         | Topic pubsub                                | my-topic        |
| GOOGLE_APPLICATION_CREDENTIALS   | Google cloud credentias                     |                 |

> Topic is created automatic case it is not exists

## Consumer
You can start the consumer with environments variables:

```
shadowd -consume
```

|Enviroment Variable               | Description                                 | Example         |
|----------------------------------|---------------------------------------------|-----------------|
| URL                              | URL target to redirect traffic              | http://myurl    |
| PROJECT_ID                       | Google cloud PROJECT_ID                     | my-project      |
| TOPIC_ID                         | Topic pubsub                                | my-topic        |
| SUBSCRIPTION_ID                  | Subscription pubsub                         | my-sub          |
| GOOGLE_APPLICATION_CREDENTIALS   | Google cloud credentias                     |                 |


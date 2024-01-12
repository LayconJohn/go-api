package akafka

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "go-api",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	kafkaConsumer.SubscribeTopics(topics, nil)
	for {
		msg, err := kafkaConsumer.ReadMessages(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}

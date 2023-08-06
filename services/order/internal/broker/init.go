package broker

import (
	"context"
	"os"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

func RunConsumers(ctx context.Context, handlers map[string]sarama.ConsumerGroupHandler) {
	kafkaConsumerGroups := initAllConsumerGroups()

	for topic, group := range kafkaConsumerGroups {
		go func(topic string, group *sarama.ConsumerGroup) {
			defer func() {
				if r := recover(); r != nil {
					logrus.Fatalf("kafka init topic: %s", r)
				}
			}()

			for {
				err := (*group).Consume(ctx, []string{topic}, handlers[topic])
				if err != nil {
					logrus.Fatalf("consumer group error")
				}
			}
		}(topic, group)
	}
}

func initAllConsumerGroups() map[string]*sarama.ConsumerGroup {
	return map[string]*sarama.ConsumerGroup{
		os.Getenv("PAYMENT_STATUS_TOPIC"):  initGroup(os.Getenv("PAYMENT_STATUS_TOPIC")),
		os.Getenv("DELIVERY_STATUS_TOPIC"): initGroup(os.Getenv("DELIVERY_STATUS_TOPIC")),
	}
}

func initGroup(topic string) *sarama.ConsumerGroup {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V2_3_0_0
	cfg.Consumer.Return.Errors = true

	group, err := sarama.NewConsumerGroup([]string{os.Getenv("KAFKA_ADDR")}, topic, cfg)
	if err != nil {
		logrus.Fatalf("Message hasn't been marshaled.")
		return nil
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logrus.Fatalf("kafka init group error: %s", r)
			}
		}()

		for err := range group.Errors() {
			logrus.Fatalf("consumer group error: %s", err)
		}
	}()

	return &group
}

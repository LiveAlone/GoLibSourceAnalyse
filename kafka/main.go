package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

var (
	broker   = "123.156.229.185:9092"
	username = "admin"
	password = "education@4"
	topic    = "shengxue_role"
)

func main() {
	v, err := sarama.ParseKafkaVersion("0.11.0.3")
	if err != nil {
		panic("Error parsing Kafka version: " + err.Error())
	}

	config := sarama.NewConfig()
	config.Net.SASL.Enable = true
	config.Net.SASL.User = username
	config.Net.SASL.Password = password
	config.Version = v
	//client, err := sarama.NewClient([]string{broker}, config)
	//if err != nil {
	//	log.Fatalf("unable to create kafka client: %q", err)
	//}
	gp, err := sarama.NewConsumerGroup([]string{broker}, "zyb_shengxue_role", config)
	if err != nil {
		log.Fatal(err)
	}
	defer gp.Close()

	gp.Consume(context.TODO(), []string{topic}, &ConsumerGroup{})
}

type ConsumerGroup struct {
}

func (c *ConsumerGroup) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *ConsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Println(string(message.Value))
		session.MarkMessage(message, "")
	}
	return nil
}

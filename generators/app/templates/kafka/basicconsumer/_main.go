package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	logger "github.com/ricardo-ch/go-logger"
	"<%- repourl%>/config"
)


func init() {
	
	// initialization (optional)
	logger.InitLogger(false)
}

func main() {

	configKafka := cluster.NewConfig()
	configKafka.Consumer.Return.Errors = true
	configKafka.Consumer.Offsets.Initial = sarama.OffsetOldest
	configKafka.Version = sarama.V1_0_0_0

	consumer, err := cluster.NewConsumer(strings.Split(config.KafkaBrokers, ","), "YOUR_GROUP_ID", []string{"YOUR_TOPIC", "YOUR_TOPIC2"}, configKafka)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case err := <-consumer.Errors():
			fmt.Println(err, "consumer.Errors()")
		case msg := <-consumer.Messages():
			fmt.Println("---------------------------------------------------")
			fmt.Println("TOPIC= ", msg.Topic)
			fmt.Println("Message= ", string(msg.Value))
			fmt.Println("Timestamp = ", msg.Timestamp)
			consumer.MarkOffset(msg, "")
		case <-signals:
			fmt.Println("Interruption is detected")
			os.Exit(1)
		}
	}
}

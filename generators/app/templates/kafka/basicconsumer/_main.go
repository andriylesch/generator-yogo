package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/Shopify/sarama"
	logger "github.com/ricardo-ch/go-logger"
	"<%- repourl%>/config"
)


func init() {
	
	// initialization (optional)
	logger.InitLogger(false)
}

func main() {
	kconfig := sarama.NewConfig()
	kconfig.Version = sarama.V1_0_0_0

	consumer, err := sarama.NewConsumer(strings.Split(config.KafkaBrokers, ","), kconfig)
	if err != nil {
		logger.Error(err.Error())
	}

	defer func() {
		if localErr := consumer.Close(); localErr != nil {
			logger.Error(localErr.Error())
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("YOUR_TOPIC", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			logger.Error(err.Error())
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Println("----- Consumed message ----")
			fmt.Printf("Time 	 : %s\n", msg.Timestamp)
			fmt.Printf("Key 	 : %s\n", msg.Key)
			fmt.Printf("Message : %s\n", msg.Value)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	logger.Info(fmt.Sprintf("Consumed: %d\n", consumed))
}

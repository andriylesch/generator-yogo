package producer

import (
	"time"

	"github.com/Shopify/sarama"
)

type Producer struct {
	Sync sarama.SyncProducer
}

var newSyncProducer = sarama.NewSyncProducer

func NewProducer(config *sarama.Config, brokers []string) (Producer, error) {

	if config == nil {
		config := sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForLocal
		config.Producer.Retry.Max = 10
		config.Producer.Return.Successes = true
		config.Version = sarama.V1_0_0_0
	}

	sync, err := newSyncProducer(brokers, config)
	if err != nil {
		return Producer{
			Sync: sync,
		}, err
	}

	return Producer{
		Sync: sync,
	}, nil
}

// ProduceMessage : produces the kafka message
func (p Producer) ProduceMessage(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(message),
		Timestamp: time.Now().UTC(),
	}

	_, _, err := p.Sync.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

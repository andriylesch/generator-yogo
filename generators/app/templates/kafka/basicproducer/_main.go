package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	logger "github.com/ricardo-ch/go-logger"
	"<%- repourl%>/producer"
	"<%- repourl%>/config"
	uuid "github.com/satori/go.uuid"
)

// User ...
type User struct {
	UserID    string `json:"user_id"`
	Nick      string `json:"nick"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func init() {

	// initialization (optional)
	logger.InitLogger(false)

}

func main() {

	user := User{
		UserID:    uuid.NewV4().String(),
		Nick:      "test_nick",
		Email:     "test@gmail.com",
		FirstName: "FN",
		LastName:  "LN",
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// init Producer
	kconfig := sarama.NewConfig()
	kconfig.Producer.RequiredAcks = sarama.WaitForLocal
	kconfig.Producer.Retry.Max = 10
	kconfig.Producer.Return.Successes = true
	kconfig.Version = sarama.V1_0_0_0

	prod, err := producer.NewProducer(kconfig, strings.Split(config.KafkaBrokers, ","))
	if err != nil {
		logger.Error(fmt.Sprintf("[ERROR] Failed to start Sarama producer: %s\n", err.Error()))
	}

	prod.ProduceMessage("YOUR_TOPIC", sarama.ByteEncoder(userJSON))

	fmt.Println("Message : ", string(userJSON))
}


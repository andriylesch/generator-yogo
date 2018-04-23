package main

import (
	"encoding/json"
	"fmt"
	"strings"

	logger "github.com/ricardo-ch/go-logger"
	"<%- repourl%>/producer"
	"<%- repourl%>/config"
	uuid "github.com/satori/go.uuid"
)


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
		FirstName: "firstName",
		LastName:  "lastName",
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// kconfig := sarama.NewConfig()
	// kconfig.Version = sarama.V1_0_0_0
	// kconfig.Producer.Return.Errors = true

	// init Producer
	producer, err := producer.NewProducer(nil, strings.Split(config.KafkaBrokers, ","))
	if err != nil {
		logger.Error(err.Error())
		return
	}

	producer.ProduceMessage("YOUR_TOPIC", userJSON)
	fmt.Println("test message: ", string(userJSON))
}

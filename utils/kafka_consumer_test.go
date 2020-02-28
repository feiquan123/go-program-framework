package utils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// TempKafkaMsg : kafka msg struct , topic : "topic_temp"
type TempKafkaMsg struct {
	Word string `json:"word"`
}

func (t *TempKafkaMsg) ParseKafkaMsg(message *sarama.ConsumerMessage) (err error) {
	msg := string(message.Value)
	if msg == "" {
		return nil
	}

	if err = t.jsonParse(msg); err != nil {
		fmt.Println("jsonParse err:", err)
		return err
	}
	return nil
}

func (t *TempKafkaMsg) jsonParse(msg string) (err error) {
	err = json.Unmarshal([]byte(msg), t)
	if err != nil {
		return err
	}
	return
}

func TestMultAssocComsumer(t *testing.T) {
	// kafka_topic_driver_action
	var (
		Topic    = "topic_temp"      // topic
		Consumer = "topic_temp_test" // Consumer group
		logger   = logrus.New()
	)
	MultAssocComsumer(
		23,
		[]string{"127.0.0.1:9092"},
		Consumer,
		Topic,
		func(topic string) (msg Msg) {
			if topic == Topic {
				msg = new(TempKafkaMsg)
			} else {
				logger.Panic(topic + " is unknown topic.")
			}
			return
		},
		logger,
		10000,
		10000,
	)
}

package services

import (
	"fmt"
	"log"

	"github.com/GodCratos/retail_client/configs"
	"github.com/nsqio/go-nsq"
)

func NsqPushMessage(message []byte, topicName string) error {
	fmt.Println("----------------------------------------------------------------------------")
	log.Println("[SERVICES:QUEUE:PRODUCER] Start push message")
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(configs.NsqGetUrl(), config)
	if err != nil {
		log.Println("[SERVICES:QUEUE:PRODUCER] Could not create producer. Error description : ", err)
		return err
	}
	log.Println("[SERVICES:QUEUE:PRODUCER] Producer created")

	err = producer.Publish(topicName, message)
	if err != nil {
		log.Println("[SERVICES:QUEUE:PRODUCER] Could not send message. Error description : ", err)
		return err
	}
	log.Println("[SERVICES:QUEUE:PRODUCER] Message pushed : ", string(message))
	producer.Stop()
	return nil
}

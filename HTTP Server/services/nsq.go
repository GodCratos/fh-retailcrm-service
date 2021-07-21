package services

import (
	"fmt"
	"log"

	"github.com/GodCratos/mindbox/configs"
	"github.com/nsqio/go-nsq"
)

func NsqPushMessage(message []byte, topicName string) error {
	fmt.Println("----------------------------------------------------------------------------")
	log.Println("[SERVICES:NSQ] Start push message")
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(configs.NsqGetUrl(), config)
	if err != nil {
		log.Println("[SERVICES:NSQ] Could not create producer. Error description : ", err)
		return err
	}
	log.Println("[SERVICES:NSQ] Producer created")

	err = producer.Publish(topicName, message)
	if err != nil {
		log.Println("[SERVICES:NSQ] Could not send message. Error description : ", err)
		return err
	}
	log.Println("[SERVICES:NSQ] Message pushed : ", string(message))
	producer.Stop()
	return nil
}

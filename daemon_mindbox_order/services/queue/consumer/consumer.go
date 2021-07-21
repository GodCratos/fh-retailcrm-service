package queue

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/GodCratos/mindbox_order/configs"
	"github.com/GodCratos/mindbox_order/services"
	"github.com/GodCratos/mindbox_order/structures"
	"github.com/nsqio/go-nsq"
)

type messageHandler struct{}

func NsqListen() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(configs.NsqGetTopicMindboxOrder(), configs.NsqGetChannel(), config)
	if err != nil {
		log.Panic("[SERVICES:QUEUE:CONSUMER] Could not create consumer. Error description : ", err)
	}
	consumer.AddHandler(&messageHandler{})
	err = consumer.ConnectToNSQD(configs.NsqGetUrl())
	if err != nil {
		log.Panic("[SERVICES:QUEUE:CONSUMER] Could not connect. Error description : ", err)
	}
	wg.Wait()
}

func (h *messageHandler) HandleMessage(message *nsq.Message) error {
	fmt.Println("----------------------------------------------------------------------------")
	log.Println("[SERVICES:QUEUE:CONSUMER] Message received. Message : ", string(message.Body))
	if len(message.Body) == 0 {
		log.Println("[SERVICES:QUEUE:CONSUMER] Message is empty")
		return nil
	}
	var QueueMessage structures.MessageFromNSQ
	err := json.Unmarshal(message.Body, &QueueMessage)
	if err != nil {
		log.Println("[SERVICES:QUEUE:CONSUMER] Wrong message. Error description : ", err)
		return err
	}
	err = services.MindboxSendRequest(QueueMessage)
	if err != nil {
		log.Println("[SERVICES:QUEUE:CONSUMER] Error while getting client . Error description : ", err)
		services.SentryPushMessage(err.Error())
		return err
	}
	return nil
}

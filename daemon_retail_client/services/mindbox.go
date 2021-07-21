package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/GodCratos/retail_client/configs"
	producer "github.com/GodCratos/retail_client/services/queue/producer"
	"github.com/GodCratos/retail_client/structures"
)

const (
	layout = "2006-01-02 15:04:05"
)

func MindboxCreateStructureToSend(retailStruct map[string]interface{}, itsNewClient bool) error {
	customer := retailStruct["customers"].([]interface{})[0].(map[string]interface{})
	var strMindboxClient structures.MindboxClient
	strMindboxClient.Customer.Ids.RetailCRMID = strconv.FormatFloat(customer["id"].(float64), 'f', 0, 64)
	if value, ok := customer["externalId"]; ok {
		strMindboxClient.Customer.Ids.WebsiteID = value.(string)
	}
	strMindboxClient.Customer.FullName = customer["firstName"].(string)
	if value, ok := customer["lastName"]; ok {
		strMindboxClient.Customer.FullName = value.(string) + " " + strMindboxClient.Customer.FullName
	}
	if value, ok := customer["patronymic"]; ok {
		strMindboxClient.Customer.FullName = strMindboxClient.Customer.FullName + " " + value.(string)
	}
	if value, ok := customer["email"]; ok {
		strMindboxClient.Customer.Email = value.(string)
	}

	if len(customer["phones"].([]interface{})) != 0 {
		strMindboxClient.Customer.MobilePhone = customer["phones"].([]interface{})[0].(map[string]interface{})["number"].(string)
	}
	strMindboxClient.Customer.CustomFields.Address = ""
	if _, ok := customer["address"]; ok {
		if value, ok := customer["address"].(map[string]interface{})["text"]; ok {
			strMindboxClient.Customer.CustomFields.Address = value.(string)
		}
	}

	if itsNewClient {
		strMindboxClient.Customer.Subscriptions = append(strMindboxClient.Customer.Subscriptions, struct {
			IsSubscribed bool "json:\"isSubscribed\""
		}{})
		strMindboxClient.Customer.Subscriptions[0].IsSubscribed = true
	}

	if value, ok := customer["birthday"]; ok {
		strMindboxClient.Customer.BirthDate = value.(string)
	}

	/*nowTime := time.Now()
	dur, _ := time.ParseDuration("-1m")
	t := nowTime.Add(dur)
	strMindboxClient.ExecutionDateTimeUtc = t.Format(layout)*/
	strMindboxClient.New = itsNewClient
	jsonMsg, err := json.Marshal(strMindboxClient)
	if err != nil {
		log.Println("[SERVICES:MINDBOX] Error while generating JSON. Error description : ", err)
		return errors.New("[SERVICES:MINDBOX] Error while generating JSON")
	}
	err = producer.NsqPushMessage(jsonMsg, configs.NsqGetTopicMindboxClient())
	if err != nil {
		return errors.New(fmt.Sprintf("[SERVICES:RETAIL] Ошибка при отправке сообщения в nsq. Error description : %s", err.Error()))
	}
	return nil
}

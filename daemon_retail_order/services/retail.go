package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GodCratos/retail_order/configs"
	"github.com/GodCratos/retail_order/structures"
)

func RetailGetOrderByFilter(queueMessage structures.MessageFromNSQ) error {
	clientID := queueMessage.ID
	fmt.Println("----------------------------------------------------------------------------")
	url := configs.RetailGetURL(clientID)
	log.Println("[SERVICES:RETAIL] Start sending request : ", url)
	clientGet := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("[SERVICES:RETAIL] Error while creating a new request. Error description : ", err)
		return errors.New(fmt.Sprintf("[SERVICES:RETAIL] Ошибка при создании нового запроса. Error description : %s", err.Error()))
	}
	resp, err := clientGet.Do(req)
	if err != nil {
		log.Println("[SERVICES:RETAIL] Error while sending request. Error description : ", err)
		return errors.New(fmt.Sprintf("[SERVICES:RETAIL] Ошибка при отправке запроса в Retail. Error description : %s", err.Error()))
	}
	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)
	log.Println("[SERVICES:RETAIL] Response from Retail : ", string(respByte))

	retailStruct, err := RetailParserJSON(respByte)
	if err != nil {
		return err
	}
	if len(retailStruct["orders"].([]interface{})) == 0 {
		log.Println("[SERVICES:MINDBOX] Orders array is empty")
		return nil
	}
	err = MindboxCreateStructureToSend(retailStruct, queueMessage.New)
	if err != nil {
		return err
	}
	return nil
}

func RetailParserJSON(value []byte) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(value, &jsonMap)
	if err != nil {
		log.Println("[SERVICES:RETAIL] Error while parsing JSON. Error description : ", err)
		return nil, errors.New(fmt.Sprintf("[SERVICES:RETAIL] Ошибка при разборе JSON. Error description : %s", err.Error()))
	}
	return jsonMap, nil
}

func RetailGetDescriptionTypePromoAction(nameType string) string {
	switch nameType {
	case "manual_order":
		return "Разовая скидка на заказ"
	case "manual_product":
		return "Дополнительная скидка на товар"
	case "loyalty_level":
		return "Скидка по уровню программы лояльности"
	case "loyalty_event":
		return "Скидка по событию программы лояльности"
	case "personal":
		return "Персональная скидка"
	case "bonus_charge":
		return "Списание бонусов ПЛ"
	case "round":
		return "Скидка от округления"
	}
	return "Неизвестная скидка"
}

func RetailIsStatusCancel(status string) bool {
	if status == "no-call" ||
		status == "no-product" ||
		status == "already-buyed" ||
		status == "delyvery-did-not-suit" ||
		status == "prices-did-not-suit" ||
		status == "cancel-other" ||
		status == "cancelled" ||
		status == "disbanded" ||
		status == "return" ||
		status == "partial-return" {
		return true
	}
	return false
}

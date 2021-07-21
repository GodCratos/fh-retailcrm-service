package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/GodCratos/retail_order/configs"
	producer "github.com/GodCratos/retail_order/services/queue/producer"
	"github.com/GodCratos/retail_order/structures"
)

const (
	layout = "2006-01-02 15:04:05"
)

func MindboxCreateStructureToSend(retailStruct map[string]interface{}, itsNewOrder bool) error {
	order := retailStruct["orders"].([]interface{})[0].(map[string]interface{})
	var strMindboxOrder structures.MindboxOrder
	if value, ok := order["customer"].(map[string]interface{})["externalId"]; ok {
		strMindboxOrder.Customer.IDs.WebsiteID = value.(string)
	} else {
		strMindboxOrder.Customer.IDs.WebsiteID = fmt.Sprintf("%v", order["customer"].(map[string]interface{})["id"].(float64))
	}
	strMindboxOrder.Customer.IDs.RetailCRMID = fmt.Sprintf("%v", order["customer"].(map[string]interface{})["id"].(float64))
	if value, ok := order["externalId"]; ok {
		strMindboxOrder.Order.IDs.WebsiteID = value.(string)
	} else {
		strMindboxOrder.Order.IDs.WebsiteID = fmt.Sprintf("%v", order["id"].(float64))
	}
	if value, ok := order["delivery"]; ok {
		strMindboxOrder.Order.DeliveryCost = value.(map[string]interface{})["cost"].(float64)
		if _, ok := value.(map[string]interface{})["code"]; ok {
			strMindboxOrder.Order.CustomFields.DeliveryType = value.(map[string]interface{})["code"].(string)
		}

		if _, ok := value.(map[string]interface{})["address"]; ok {
			if _, ok := value.(map[string]interface{})["address"].(map[string]interface{})["text"]; ok {
				strMindboxOrder.Order.CustomFields.DeliveryAddress = value.(map[string]interface{})["address"].(map[string]interface{})["text"].(string)
			}
		}

		if _, ok := value.(map[string]interface{})["date"]; ok {
			strMindboxOrder.Order.CustomFields.DeliveryDate = value.(map[string]interface{})["date"].(string)
		}

		if _, ok := value.(map[string]interface{})["time"]; ok {
			strMindboxOrder.Order.CustomFields.DeliveryTime = value.(map[string]interface{})["time"].(map[string]interface{})["from"].(string) + "-" +
				value.(map[string]interface{})["time"].(map[string]interface{})["to"].(string)
		}
	}
	if _, ok := order["firstName"]; ok {
		strMindboxOrder.Order.CustomFields.Name = order["firstName"].(string)
	}

	strMindboxOrder.Order.CustomFields.Status = order["status"].(string)
	strMindboxOrder.Order.TotalPrice = order["totalSumm"].(float64)
	if len(order["payments"].(map[string]interface{})) != 0 {
		indexPay := 0
		for _, valuePay := range order["payments"].(map[string]interface{}) {
			strMindboxOrder.Order.Payments = append(strMindboxOrder.Order.Payments, struct {
				Type   string  "json:\"type,omitempty\""
				Amount float64 "json:\"amount,omitempty\""
			}{})
			strMindboxOrder.Order.Payments[indexPay].Type = valuePay.(map[string]interface{})["type"].(string)
			strMindboxOrder.Order.Payments[indexPay].Amount = valuePay.(map[string]interface{})["amount"].(float64)
			indexPay++
		}
	}
	if order["summ"].(float64) == 0 {
		strMindboxOrder.Order.Payments = nil
	}

	for index, value := range order["items"].([]interface{}) {
		item := value.(map[string]interface{})
		strMindboxOrder.Order.Lines = append(strMindboxOrder.Order.Lines, struct {
			MinPricePerItem        float64     "json:\"minPricePerItem,omitempty\""
			CostPricePerItem       float64     "json:\"costPricePerItem,omitempty\""
			BasePricePerItem       float64     "json:\"basePricePerItem,omitempty\""
			Quantity               interface{} "json:\"quantity,omitempty\""
			QuantityType           string      "json:\"quantityType,omitempty\""
			DiscountedPricePerLine float64     "json:\"discountedPricePerLine,omitempty\""
			LineID                 string      "json:\"lineId,omitempty\""
			LineNumber             int         "json:\"lineNumber,omitempty\""
			Discounts              []struct {
				Type                string "json:\"type,omitempty\""
				ExternalPromoAction struct {
					IDs struct {
						ExternalId string "json:\"externalId,omitempty\""
					} "json:\"ids,omitempty\""
				} "json:\"externalPromoAction,omitempty\""
				Amount float64 "json:\"amount,omitempty\""
			} "json:\"discounts,omitempty\""
			Product struct {
				IDs struct {
					Website string "json:\"website,omitempty\""
					ERP     string "json:\"eRP,omitempty\""
				} "json:\"ids,omitempty\""
			} "json:\"product,omitempty\""
			Status string "json:\"status,omitempty\""
		}{})

		if value, ok := item["discountTotal"]; ok {
			strMindboxOrder.Order.Lines[index].MinPricePerItem = item["initialPrice"].(float64) - value.(float64)
			strMindboxOrder.Order.Lines[index].DiscountedPricePerLine = (item["initialPrice"].(float64) - value.(float64)) * item["quantity"].(float64)
		} else {
			strMindboxOrder.Order.Lines[index].MinPricePerItem = item["initialPrice"].(float64)
			strMindboxOrder.Order.Lines[index].DiscountedPricePerLine = item["initialPrice"].(float64) * item["quantity"].(float64)
		}
		strMindboxOrder.Order.Lines[index].CostPricePerItem = item["initialPrice"].(float64)
		strMindboxOrder.Order.Lines[index].BasePricePerItem = item["initialPrice"].(float64)
		if reflect.ValueOf(item["quantity"]).Kind() == reflect.Float64 {
			strMindboxOrder.Order.Lines[index].Quantity = item["quantity"].(float64)
			strMindboxOrder.Order.Lines[index].QuantityType = "double"
		} else {
			strMindboxOrder.Order.Lines[index].Quantity = item["quantity"].(int)
			strMindboxOrder.Order.Lines[index].QuantityType = "int"
		}
		strMindboxOrder.Order.Lines[index].LineID = fmt.Sprintf("%v", item["id"].(float64))
		strMindboxOrder.Order.Lines[index].LineNumber = index + 1

		if len(item["discounts"].([]interface{})) != 0 {
			for indexDis, valueDis := range item["discounts"].([]interface{}) {
				strMindboxOrder.Order.Lines[index].Discounts = append(strMindboxOrder.Order.Lines[index].Discounts, struct {
					Type                string "json:\"type,omitempty\""
					ExternalPromoAction struct {
						IDs struct {
							ExternalId string "json:\"externalId,omitempty\""
						} "json:\"ids,omitempty\""
					} "json:\"externalPromoAction,omitempty\""
					Amount float64 "json:\"amount,omitempty\""
				}{})
				if valueType, ok := valueDis.(map[string]interface{})["type"]; ok {
					strMindboxOrder.Order.Lines[index].Discounts[indexDis].Type = "externalPromoAction"
					strMindboxOrder.Order.Lines[index].Discounts[indexDis].ExternalPromoAction.IDs.ExternalId = RetailGetDescriptionTypePromoAction(valueType.(string))
				}
				strMindboxOrder.Order.Lines[index].Discounts[indexDis].Amount = valueDis.(map[string]interface{})["amount"].(float64)
			}
		}
		/*if value, ok := item["offer"].(map[string]interface{})["externalId"]; ok {
			strMindboxOrder.Order.Lines[index].Product.IDs.Website = value.(string)
		}*/

		if value, ok := item["offer"].(map[string]interface{})["xmlId"]; ok {
			strMindboxOrder.Order.Lines[index].Product.IDs.ERP = value.(string)
		}

		if !itsNewOrder {
			strMindboxOrder.Order.Lines[index].Status = item["status"].(string)
		}

	}
	if value, ok := order["email"]; ok {
		strMindboxOrder.Order.Email = value.(string)
	}
	if value, ok := order["phone"]; ok {
		strMindboxOrder.Order.MobilePhone = value.(string)
	}
	/*t := time.Now()
	strMindboxOrder.ExecutionDateTimeUtc = t.Format(layout)*/

	if _, ok := order["orderMethod"]; ok {
		if order["orderMethod"].(string) == "offline" {
			strMindboxOrder.PointOfContact = order["shipmentStore"].(string)
		} else if order["orderMethod"].(string) == "shopping-cart" {
			strMindboxOrder.PointOfContact = "Site"
		} else if order["orderMethod"].(string) == "app" {
			strMindboxOrder.PointOfContact = "MobileApp"
		} else {
			strMindboxOrder.PointOfContact = order["orderMethod"].(string)
		}
		if _, ok := order["shipmentStore"]; ok {
			strMindboxOrder.Order.CustomFields.ShipmentStore = order["shipmentStore"].(string)
		}
	}
	if itsNewOrder {
		fullname := order["firstName"].(string)
		if _, ok := order["lastName"]; ok {
			fullname = order["lastName"].(string) + " " + fullname
		}
		strMindboxOrder.Customer.FullName = fullname
		strMindboxOrder.Customer.MobilePhone = order["phone"].(string)
		if _, ok := order["email"]; ok {
			strMindboxOrder.Customer.Email = order["email"].(string)
		}
	}
	strMindboxOrder.New = itsNewOrder
	jsonMsg, err := json.Marshal(strMindboxOrder)
	if err != nil {
		log.Println("[SERVICES:MINDBOX] Error while generating JSON. Error description : ", err)
		return errors.New("[SERVICES:MINDBOX] Error while generating JSON")
	}
	err = producer.NsqPushMessage(jsonMsg, configs.NsqGetTopicMindboxOrder())
	if err != nil {
		return errors.New(fmt.Sprintf("[SERVICES:RETAIL] Ошибка при отправке сообщения в nsq. Error description : %s", err.Error()))
	}
	return nil
}

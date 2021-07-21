package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/GodCratos/mindbox/configs"
	"github.com/GodCratos/mindbox/services"
	"github.com/GodCratos/mindbox/structures"
)

func HandlerOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	new := r.URL.Query().Get("new")
	if id == "" || new == "" {
		log.Println("[HANDLER:ORDER] Parameters are empty")
		return
	}
	var message structures.Message
	message.ID = id
	message.New = false
	if new == "yes" {
		message.New = true
	}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Println("[HANDLER:ORDER] Error when marshalling json. Error description : ", err)
		services.RetailCreateErrorResponse(w, errors.New("Ошибка при формировании JSON").Error())
		services.SentryPushMessage(fmt.Sprintf("[HANDLER:ORDER] Ошибка при формировании JSON. Error description : %s", err.Error()))
		return
	}
	err = services.NsqPushMessage(jsonMessage, configs.NsqGetTopicRetailOrder())
	if err != nil {
		services.RetailCreateErrorResponse(w, errors.New("Ошибка при отправке сообщения в nsq").Error())
		services.SentryPushMessage(fmt.Sprintf("[HANDLER:ORDER] Ошибка при отправке сообщения в nsq. Error description : %s", err.Error()))
		return
	}

	var resp structures.Response
	resp.Success = true
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Println("[HANDLER:ORDER] Error parsing JSON. Error description : ", err)
		services.SentryPushMessage(fmt.Sprintf("[HANDLER:ORDER] Ошибка при формировании JSON. Error description : %s", err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

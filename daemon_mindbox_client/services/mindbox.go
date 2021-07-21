package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GodCratos/mindbox_client/configs"
	"github.com/GodCratos/mindbox_client/structures"
)

func MindboxSendRequest(strRequest structures.MessageFromNSQ) error {
	url := configs.MindboxGetUrl(strRequest.New)
	strRequest.New = false
	jsonStr, err := json.Marshal(strRequest)
	if err != nil {
		log.Println("[SERVICES:MINDBOX] Error while generating JSON. Error description : ", err)
		return errors.New("[SERVICES:MINDBOX] Error while generating JSON")
	}
	reqBody := bytes.NewReader(jsonStr)
	fmt.Println("----------------------------------------------------------------------------")
	log.Println("[SERVICES:MINDBOX] Start sending request : ", url)
	log.Println("[SERVICES:MINDBOX] Request body : ", string(jsonStr))
	clientPost := http.Client{}
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		log.Println("[SERVICES:MINDBOX] Error while creating a new request. Error description : ", err)
		return errors.New(fmt.Sprintf("[SERVICES:MINDBOX] Ошибка при создании нового запроса. Error description : %s", err.Error()))

	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Mindbox secretKey=\"%s\"", configs.MindboxGetSecretKey()))
	resp, err := clientPost.Do(req)
	if err != nil {
		log.Println("[SERVICES:MINDBOX] Error while sending request. Error description : ", err)
		return errors.New(fmt.Sprintf("[SERVICES:MINDBOX] Ошибка при отправке запроса в Mindbox. Error description : %s", err.Error()))
	}
	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)
	log.Println("[SERVICES:MINDBOX] Response from Mindbox : ", string(respByte))
	mindboxStruct, err := MindboxParserJSON(respByte)
	if err != nil {
		return err
	}
	if value, ok := mindboxStruct["errorMessage"]; ok {
		return errors.New(value.(string))
	}
	return nil
}

func MindboxParserJSON(value []byte) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(value, &jsonMap)
	if err != nil {
		log.Println("[SERVICES:MINDBOX] Error while parsing JSON. Error description : ", err)
		return nil, errors.New(fmt.Sprintf("[SERVICES:MINDBOX] Ошибка при разборе JSON. Error description : %s", err.Error()))
	}
	return jsonMap, nil
}

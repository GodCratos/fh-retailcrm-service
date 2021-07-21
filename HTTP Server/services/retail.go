package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GodCratos/mindbox/structures"
)

func RetailCreateErrorResponse(w http.ResponseWriter, message string) {
	var errorMsg structures.ResponseError
	errorMsg.Success = false
	errorMsg.ErrorMsg = message
	req, err := json.Marshal(errorMsg)
	if err != nil {
		log.Println("[SERVICES:RETAIL] Error parsing JSON. Error description : ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(req)
}

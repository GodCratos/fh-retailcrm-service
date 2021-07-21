package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/GodCratos/mindbox/configs"
	"github.com/GodCratos/mindbox/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/order", handlers.HandlerOrder).Methods("GET")
	router.HandleFunc("/client", handlers.HandlerClient).Methods("GET")
	http.Handle("/", router)
	log.Println(fmt.Sprintf("[MAIN] HTTP server is running on port %s", configs.HTTPGetPort()))
	go http.ListenAndServe(configs.HTTPGetPort(), nil)
	runtime.Goexit()
}

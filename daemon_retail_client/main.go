package main

import (
	"runtime"

	consumer "github.com/GodCratos/retail_client/services/queue/consumer"
)

func main() {
	go consumer.NsqListen()
	runtime.Goexit()
}

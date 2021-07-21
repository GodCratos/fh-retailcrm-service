package main

import (
	"runtime"

	consumer "github.com/GodCratos/mindbox_client/services/queue/consumer"
)

func main() {
	go consumer.NsqListen()
	runtime.Goexit()
}

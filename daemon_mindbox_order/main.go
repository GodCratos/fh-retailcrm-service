package main

import (
	"runtime"

	consumer "github.com/GodCratos/mindbox_order/services/queue/consumer"
)

func main() {
	go consumer.NsqListen()
	runtime.Goexit()
}

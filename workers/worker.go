package main

import (
	"asynq_demo/common"
)

func main() {
	for i := 0; i < 100; i++ {
		common.EmailDeliveryTaskAdd(i)
		//time.Sleep(time.Second * 3)
	}
}

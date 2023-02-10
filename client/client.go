package main

import (
	"asynq_demo/common"
	"github.com/hibiken/asynq"
	"log"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       2,
		},
		asynq.Config{
			// 每个进程并发执行的worker数量
			Concurrency: 5,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(common.TypeEmailDelivery, common.HandleEmailDeliveryTask)
	mux.HandleFunc(common.TypeEcho, common.HandleEchoTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

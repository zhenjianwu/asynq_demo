package main

import (
	"asynq_demo/common"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

func main() {

	// 创建任务
	task, _ := common.NewEchoTask(42, fmt.Sprintf("some:template:id:%d", 10), `{"name":"lisi"}`)

	// 创建调度器
	scheduler := asynq.NewScheduler(asynq.RedisClientOpt{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       2,
	}, &asynq.SchedulerOpts{})

	// 注册任务
	// 可以使用 cron 规范字符串来指定计划。
	_, _ = scheduler.Register("* * * * *", task) // 每分钟执行一次
	// 也可以使用"@every <duration>"语法指定间隔
	//entryID, err = scheduler.Register("@every 30s", task) // 每30s执行一次
	//// 可以在注册任务的同时，指定配置项
	//entryID, err = scheduler.Register("@every 24h", task, asynq.Queue("myqueue")) // 每24小时执行一次 队列名字“myqueue”。

	// 运行调度器 并阻塞
	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}

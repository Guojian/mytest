package scheduler

import (
	"fmt"
	"leader"
	"common"
	"follower"
)

type TaskScheduler struct {
	mq_handler *common.MQHandler
	mq_channel chan common.Message
}

func InitTaskScheduler(poolSize int) *TaskScheduler {
	scheduler := new(TaskScheduler)
	mq_channel := make(chan common.Message, 10)
	scheduler.mq_channel = mq_channel
	mq_handler := common.InitMQHandler(10, mq_channel)
	scheduler.mq_handler = mq_handler
	return scheduler
}


func (scheduler TaskScheduler) Loop() int {
	go scheduler.mq_handler.Loop()
	for msg := range scheduler.mq_channel{
		switch task:=msg.GetTask(); task {
			case "Leader":
				go leader.Handle(msg.GetRelatedUuid())
			case "Follower":
				go follower.Handle("1", "b")
			default:
				fmt.Println("Handler not found")
		}
	}
	return 1
}
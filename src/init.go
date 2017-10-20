package main

import (
	"common"
	"resources"
	"scheduler"
	"fmt"
)


func main() {
	
	vm := resources.NewVM()
	abc := vm.Init("CPU", "Memory")
	disk := common.NewDisk()

	var res []interface{}
	res = append(res, abc.Dumps())
	res = append(res, disk)
	//res[0] = vm
	//res[1] = disk


	job := common.NewJob("myQueue", res)

	fmt.Println(res[0])
	//job.Resources = append(job.Resources, vm)
	//job.Resources = append(job.Resources, disk)
	for i:=1; i < 5; i++{
		var (
			task string = "Leader"
			args = []string {"a", "b", "c"}
			queue string = "192.168.1.1"
			relatedUuid string = job.GetId()
			priority int = 1
		)

		msg := common.NewMessage(task, args, queue, relatedUuid, priority)
		common.MongoSet(msg)
	}

	for i:=1; i < 5; i++{
		var (
			task string = "Follower"
			args = []string {"a", "b", "c"}
			queue string = "192.168.1.1"
			relatedUuid string = job.GetId()
			priority int = 1
		)

		msg := common.NewMessage(task, args, queue, relatedUuid, priority)
		common.MongoSet(msg)
	}
	

    taskScheduler := scheduler.InitTaskScheduler(10)
    taskScheduler.Loop()

}

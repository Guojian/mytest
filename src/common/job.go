package common

import (
	"github.com/satori/go.uuid"
	"fmt"
)

type Job interface{
	GetId() string
	Empty() job
}

type job struct {
	JobId string
	TaskList []string
	Queue string
	CreateTime int
	FinishTime int
	Type string
	State string
	Description string
	Handler string
	Resources []interface{}
}

func (thisJob job) GetId() string {
	return thisJob.JobId
}

func (thisJob job) Empty() job {
	return job{}
}

func NewJob(queue string, resources []interface{}) Job {
	uid := uuid.NewV4().String()
	fmt.Println(uid)
	return job{JobId: uid, Queue:queue, Resources:resources}
}
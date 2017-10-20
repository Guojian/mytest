package common

import (
	"github.com/satori/go.uuid"
	"time"
)

type message struct {
	Task string 
	Args []string
	Queue string
	Uuid string
	RelatedUuid string
	Priority int
	CreateTime int64
}

type Message interface{
	Dumps() message
	Empty() message
	GetRelatedUuid() string
	GetTask() string
}

func (msg message) Dumps() message {
	return msg
}

func (msg message) Empty() message {
	return message{}
}

func (msg message) GetRelatedUuid() string {
	return msg.RelatedUuid
}

func (msg message) GetTask() string {
	return msg.Task
}

func NewMessage(task string, args [] string, queue string, relatedUuid string, priority int) Message {
	uid := uuid.NewV4().String()
	now := time.Now().Unix()
	return message{Task:task, Args: args, Queue: queue, RelatedUuid: relatedUuid, Priority: priority, Uuid: uid, CreateTime: now}
}
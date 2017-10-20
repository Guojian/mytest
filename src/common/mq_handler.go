package common

import (
	//"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type MQHandler struct {
	client mgo.Collection
	mq_channel chan Message
}

func InitMQHandler(poolSize int, channel chan Message) *MQHandler {
	handler := new(MQHandler)
	session := MongoClient
	handler.client = *session.DB("test").C("message")
	handler.mq_channel = channel
	return handler
}


func (mq MQHandler) getMessage() (Message, error) {
    change := mgo.Change{Remove: true}
    //var task, args, queue, relatedUuid, priority string, string, string, string, int
    var (
		task string
		args [] string
		queue string
		relatedUuid string
		priority int
	)
    doc:= NewMessage(task, args, queue, relatedUuid, priority).Empty()
	if _, err := mq.client.Find(bson.M{}).Apply(change, &doc); err != nil {
        //panic(fmt.Errorf("get counter failed:", err.Error()))
        log.Println("Message Not Found")
        return doc, err
    }
    return doc, nil
}


func (mq MQHandler) Loop() int {
	for i:=0; i< 10; i++{
		if msg, err := mq.getMessage(); err != nil {
			time.Sleep(1 * time.Second)
		}else{
			mq.mq_channel <- msg
		}
	}
	close(mq.mq_channel)
	return 1
}
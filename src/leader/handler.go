package leader

import (
	"fmt"
	"common"
	"gopkg.in/mgo.v2/bson"
	"log"
)


func Handle(jobId string) bool{
	collection := common.MongoClient.DB("jobs").C("job")
	var res []interface{}
	job := common.NewJob("myQueue", res).Empty()

	if err := collection.Find(bson.M{}).One(&job); err != nil {
        //panic(fmt.Errorf("get counter failed:", err.Error()))
        log.Println("job Not Found")
    }
    fmt.Println(job)
	fmt.Println("leader handle done")
	return true
}
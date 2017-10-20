package common

import (
	"log"
    "gopkg.in/mgo.v2"
)


func MongoSet(obj interface{}) bool {
	session, err := mgo.Dial("mongodb://192.168.57.20:27017,192.168.57.21:27017,192.168.57.22:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("test").C("message")
    err = c.Insert(&obj)
    if err != nil {
            log.Fatal(err)
    }
    return true
}

func MongoInit() *mgo.Session {
	session, err := mgo.Dial("mongodb://192.168.57.20:27017,192.168.57.21:27017,192.168.57.22:27017")
    if err != nil {
        panic(err)
    }
    session.SetMode(mgo.Monotonic, true)
    return session
}

var MongoClient *mgo.Session = MongoInit()
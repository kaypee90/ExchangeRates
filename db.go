package main

import (
    "time"

    "gopkg.in/mgo.v2"
)

const (
	hosts = "mongodb:27017"
	database = "db"
	username = ""
	password = ""
	collection = "exchangerates"
)

type MongoStore struct {
   session *mgo.Session
}

func initialiseMongo() (session *mgo.Session){
       
	info := &mgo.DialInfo{
		Addrs: []string{hosts},
		Timeout: 60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil{
		panic(err)
	}

	return
}

func insert (mongoStore *MongoStore, exchangeRate *ExchangeRate){
	col := mongoStore.session.DB(database).C(collection)

	err := col.Insert(exchangeRate)
	if err != nil {
		panic(err)
	}

	return
}

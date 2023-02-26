package controller

import (
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

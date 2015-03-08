// paste
package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//	"html/template"
	//	"log"
	//	"os"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

type Paste struct {
	Id       bson.ObjectId
	UserId   string
	Title    string
	Content  string
	Language string
	IsPublic bool
}

type QueryPaste struct {
	UserId   string
	Title    string
	Content  string
	Language string
	IsPublic bool
}

func PublicPastes() []Paste {
	session, e := mgo.Dial("localhost")
	if e != nil {
		panic(e)
	}
	collection := session.DB("gopastebin3-3").C("pastes")
	var pastes []Paste
	err := collection.Find(bson.M{"ispublic": true}).All(&pastes)
	if err != nil {
		panic(err)
	}
	return pastes
}

func MyPastes(id string) []Paste {
	session, e := mgo.Dial("localhost")
	if e != nil {
		panic(e)
	}
	collection := session.DB("gopastebin3-3").C("pastes")
	var pastes []Paste
	err := collection.Find(bson.M{"userid": id, "ispublic": false}).All(&pastes)
	if err != nil {
		panic(err)
	}
	return pastes
}

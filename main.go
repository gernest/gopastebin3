// main
package main

import (
	"controllers"
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"net/http"
)

func init() {
	index := mgo.Index{
		Key:        []string{"ispublic"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     false,
	}
	index2 := mgo.Index{
		Key:        []string{"userid"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     false,
	}
	index3 := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	session, _ := mgo.Dial("localhost")
	c := session.DB("gopastebin3-3").C("pastes")
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	e := c.EnsureIndex(index2)
	if e != nil {
		panic(e)
	}
	e2 := c.EnsureIndex(index3)
	if e2 != nil {
		panic(e2)
	}
	http.HandleFunc("/", controllers.Index)

}

func main() {
	host := ":8888"
	fmt.Println("Opening", host+"...")

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/authorize", controllers.Authorize)
	http.HandleFunc("/oauth2callback", controllers.Callback)
	http.HandleFunc("/test", controllers.Testing)
	http.HandleFunc("/logout", controllers.Logout)
	http.HandleFunc("/new", controllers.Create)
	http.HandleFunc("/paste/", controllers.Show)

	http.ListenAndServe(host, nil)
}

// TODO: find out how to query MongoDB using Go
// TODO: add function for listing user's private pastes
// TODO: finish structs for loading sidebar
// TODO: create controller for displaying a spicific paste

// index
package controllers

import (
	//	"encoding/json"
	//	"log"
	//	jwt "github.com/dgrijalva/jwt-go"
	"html/template"
	//	"io/ioutil"
	"models"
	"net/http"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

var defaultUser = models.User{
	Id:         "0",
	Name:       "Guest",
	GivenName:  "Guest",
	FamilyName: " ",
	Link:       " ",
	Picture:    " ",
	Gender:     " ",
	Locale:     "en",
}

func Index(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Running index.go")

	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic("Template error: " + err.Error())
	}
	tmpl.Execute(w, nil)
}

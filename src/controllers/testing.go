// testing
package controllers

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"html/template"
	"io/ioutil"
	"models"
	"net/http"
	"os"
	//	"time"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

func Testing(w http.ResponseWriter, r *http.Request) {
	publicKey, e := ioutil.ReadFile("static/demo.rsa.pub")
	if e != nil {
		fmt.Println("Failure to read public key: %v", e)
		os.Exit(1)
	}
	var user models.User
	cookie, err := r.Cookie("jwtoauth1")
	if err == nil {

		tokenString := cookie.Value
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})
		obj := token.Claims["User"]

		j, _ := json.Marshal(obj)
		json.Unmarshal(j, &user)
	} else {
		user = defaultUser
	}
	/*	cookie2 := http.Cookie{Name: "jwtoauth1",
			Value:   tokenString,
			Expires: time.Now(),
		}
	http.SetCookie(w, &cookie2) */ // Keep this code around for logging off
	templ, _ := template.ParseFiles("views/layout.tpl", "views/testing.tpl")
	templ.Execute(w, struct {
		Title        string
		User         models.User
		Languages    map[string]string
		Publicpastes []models.Paste
		Mypastes     []models.Paste
	}{Title: "Add A Paste",
		User:         user,
		Languages:    models.Languages,
		Publicpastes: models.PublicPastes(),
		Mypastes:     models.MyPastes(user.Id),
	})
}

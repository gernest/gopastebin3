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
	"time"
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
	cookie, err := r.Cookie("jwtoauth1")
	if err != nil {
		fmt.Println("Cookie retrieval error:", err)
		os.Exit(2)
	}
	tokenString := cookie.Value
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	obj := token.Claims["User"]
	var user models.User
	j, _ := json.Marshal(obj)
	json.Unmarshal(j, &user)
	cookie2 := http.Cookie{Name: "jwtoauth1",
		Value:   tokenString,
		Expires: time.Now(),
	}
	http.SetCookie(w, &cookie2)
	templ, _ := template.ParseFiles("views/testing.html")
	templ.Execute(w, user.Name)
}

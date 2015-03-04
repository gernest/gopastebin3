// paste
package controllers

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"io/ioutil"
	"log"
	"models"
	"net/http"
	"os"
	"strings"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

func Create(w http.ResponseWriter, r *http.Request) {
	funcmap := template.FuncMap{
		"PublicPastes": models.PublicPastes,
	}
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
	title := r.FormValue("title")
	content := r.FormValue("content")
	language := r.FormValue("language")
	isPublicString := r.FormValue("ispublic")
	var isPublic bool
	if isPublicString == "true" || user.Id == "0" {
		isPublic = true
	} else {
		isPublic = false
	}
	userId := user.Id
	paste := models.Paste{
		Id:       bson.NewObjectId(),
		UserId:   userId,
		Title:    title,
		Content:  content,
		Language: language,
		IsPublic: isPublic,
	}
	prism := models.Languages[language]
	session, _ := mgo.Dial("localhost")
	collection := session.DB("gopastebin3-3").C("pastes")
	collection.Insert(&paste)
	log.Println("New id:", paste.Id)
	log.Println("After insert, paste =", paste)
	t, err := template.ParseFiles("views/layout.tpl", "views/create.tpl")
	if err != nil {
		log.Fatalln("template parse error:", err)
	}
	t2 := t.Funcs(funcmap)
	t2.Execute(w, struct {
		Title        string
		User         models.User
		Paste        models.Paste
		Prism        string
		Publicpastes []models.Paste
		Mypastes     []models.Paste
	}{Title: "Verify Paste",
		User:         user,
		Paste:        paste,
		Prism:        prism,
		Publicpastes: models.PublicPastes(),
		Mypastes:     models.MyPastes(user.Id),
	})
}

func Show(w http.ResponseWriter, r *http.Request) {
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
	url := r.URL.Path
	parts := strings.Split(url, "/")
	log.Println("parts[0]:", parts[0], "parts[1]:", parts[1], "parts[2]:", parts[2])
	rawId := parts[2]
	pasteId := strings.TrimLeft(rawId, "ObjectIdHex(")
	pasteId = strings.TrimRight(pasteId, ")")
	pasteId = strings.Trim(pasteId, "\"")
	realId := bson.ObjectIdHex(pasteId)
	log.Println("realId =", realId)
	session, q := mgo.Dial("localhost")
	if q != nil {
		panic(q)
	}
	collection := session.DB("gopastebin3-3").C("pastes")
	var result models.Paste
	e2 := collection.Find(bson.M{"id": realId}).One(&result)
	if e2 != nil {
		panic(e2)
	}

	prism := models.Languages[result.Language]
	t, err := template.ParseFiles("views/layout.tpl", "views/create.tpl")
	if err != nil {
		log.Fatalln("template parse error:", err)
	}
	t.Execute(w, struct {
		Title        string
		User         models.User
		Paste        models.Paste
		Prism        string
		Publicpastes []models.Paste
		Mypastes     []models.Paste
	}{Title: "Show Paste",
		User:         user,
		Paste:        result,
		Prism:        prism,
		Publicpastes: models.PublicPastes(),
		Mypastes:     models.MyPastes(user.Id),
	})

}

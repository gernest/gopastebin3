// google
package controllers

import (
	"code.google.com/p/goauth2/oauth"
	//	"config"
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"html/template"
	"io/ioutil"
	"models"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

var oauthCfg = &oauth.Config{
	//TODO: put your project's Client Id here.  To be got from https://code.google.com/apis/console
	ClientId: "821671882955-mi704nhc94u4u89jsvh6njd6ei4fug9h.apps.googleusercontent.com",

	//TODO: put your project's Client Secret value here https://code.google.com/apis/console
	ClientSecret: "WWza87_0qQ4Ed42Ckiws1RLd",

	//For Google's oauth2 authentication, use this defined URL
	AuthURL: "https://accounts.google.com/o/oauth2/auth",

	//For Google's oauth2 authentication, use this defined URL
	TokenURL: "https://accounts.google.com/o/oauth2/token",

	//To return your oauth2 code, Google will redirect the browser to this page that you have defined
	//TODO: This exact URL should also be added in your Google API console for this project within "API Access"->"Redirect URIs"
	RedirectURL: "http://localhost:8888/oauth2callback",

	//This is the 'scope' of the data that you are asking the user's permission to access. For getting user's info, this is the url that Google has defined.
	Scope: "https://www.googleapis.com/auth/userinfo.profile",
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	//Get the Google URL which shows the Authentication page to the user
	url := oauthCfg.AuthCodeURL("")

	//redirect user to that page
	http.Redirect(w, r, url, http.StatusFound)
}

const profileInfoURL = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json"

func Callback(w http.ResponseWriter, r *http.Request) {
	//Get the code from the response
	code := r.FormValue("code")

	t := &oauth.Transport{Config: oauthCfg}

	// Exchange the received code for a token
	t.Exchange(code)

	//now get user data based on the Transport which has the token
	resp, _ := t.Client().Get(profileInfoURL)

	buf := make([]byte, 1024)
	resp.Body.Read(buf)
	str := string(buf)
	str = strings.Trim(str, "\x00")
	b := []byte(str)
	//	var obj interface{}
	//	err := json.Unmarshal(b, &obj)
	var account models.User
	err := json.Unmarshal(b, &account)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		os.Exit(1)
	}
	//	m := obj.(map[string]interface{})
	//	fmt.Println("map =", m)
	/*
		account := Account{
			Id:         m["id"].(string),
			Name:       m["name"].(string),
			GivenName:  m["given_name"].(string),
			FamilyName: m["family_name"].(string),
			Link:       m["link"].(string),
			Picture:    m["picture"].(string),
			Gender:     m["gender"].(string),
			Locale:     m["locale"].(string),
		}
	*/
	//	account := obj.(Account)
	//	userInfoTemplate.Execute(w, account)
	/*	tmpl, err2 := template.ParseFiles("views/oauth2callback.html")
		if err2 != nil {
			fmt.Println("template error:", err2)
			os.Exit(2)
		}
		tmpl.Execute(w, account)
	*/
	//	CurrentUser = account
	var privateKey []byte
	privateKey, _ = ioutil.ReadFile("static/demo.rsa") // location of demo.rsa
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims["User"] = account
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	//	w.WriteHeader(http.StatusOK)
	templ, _ := template.ParseFiles("views/callback.html")
	//	r.Header.Add("token", tokenString)
	cookie := http.Cookie{Name: "jwtoauth1",
		Value:   tokenString,
		Expires: time.Now().AddDate(0, 0, 1),
	}
	http.SetCookie(w, &cookie)
	templ.Execute(w, struct{ Person models.User }{Person: account})
	//	config.LoggedIn = true
}

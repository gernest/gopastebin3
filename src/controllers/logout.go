// logout
package controllers

import (
	//	"fmt"
	"html/template"
	"net/http"
	"time"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

func Logout(w http.ResponseWriter, r *http.Request) {
	// mark the cookie for deletion
	cookie := http.Cookie{Name: "jwtoauth1",
		Value:   "logout",
		Expires: time.Now(),
	}
	http.SetCookie(w, &cookie)
	templ, _ := template.ParseFiles("views/logout.tpl")
	templ.Execute(w, nil)

}

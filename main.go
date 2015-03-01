// main
package main

import (
	"controllers"
	"fmt"
	"net/http"
)

func main() {
	host := ":8888"
	fmt.Println("Opening", host+"...")

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/authorize", controllers.Authorize)
	http.HandleFunc("/oauth2callback", controllers.Callback)
	http.HandleFunc("/test", controllers.Testing)
	http.HandleFunc("/", controllers.Index)
	http.ListenAndServe(host, nil)
}

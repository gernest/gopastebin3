// user
package models

import (
//"fmt"
)

type User struct {
	Id         string `json: "id"`
	Name       string `json: "name"`
	GivenName  string `json: "given_name"`
	FamilyName string `json: "family_name"`
	Link       string `json: "link"`
	Picture    string `json: "picture"`
	Gender     string `json: "gender"`
	Locale     string `json: "locale"`
}

/*
func main() {
	fmt.Println("Hello World!")
}
*/

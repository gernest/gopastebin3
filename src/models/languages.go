// languages
package models

import (
//	"fmt"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

type Language struct {
	Id       int
	Name     string
	Keywords []string
}

var Languages = map[string]string{"C#": "csharp", "Go": "go", "Python": "python",
	"Java": "java", "JavaScript": "javascript", "CSS": "css",
	"PHP": "php", "Ruby": "ruby", "Scala": "scala"}

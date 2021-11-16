package main

import "fmt"

const englishPrefix = "hello "
const spanishPrefix = "hola "
func main() {
	fmt.Println("hello world")
}

func getString(name string, lang string) string {
	var localName = name
	if name == "" {
		localName = "world"
	}
	if lang == "spanish" {
		return spanishPrefix + localName
	}
	return englishPrefix + localName
}
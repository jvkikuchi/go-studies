package main

import (
// "fmt"
)

var languages = map[string]string{
	"EN": "Hello, ",
	"PT": "Olá, ",
	"ES": "Hola, ",
}

func getGreeting(language string) string {
	greeting, ok := languages[language]

	if ok {
		return greeting
	}

	return "Hello, "
}

func Hello(name string, language string) string {
	greeting := getGreeting(language)

	if name == "" {
		return greeting + "World"
	}

	return greeting + name
}

func main() {
}

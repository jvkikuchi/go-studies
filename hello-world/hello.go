package main

import (
	"fmt"
)

var languages = map[string]string{
	"EN": "Hello, ",
	"PT": "Olá, ",
	"ES": "Hola, ",
}

func Hello(name string, language string) string {
	greeting, ok := languages[language]

	if (name != "" && ok) {
		return greeting + name
	}

	if (name == "" && ok) {
		return greeting + "World"
	}

	if (name == "" && !ok) {
		return "Hello, World"
	}

	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Joao", "123"))
}

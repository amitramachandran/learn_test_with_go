package main

import "fmt"

const english_prefix = "Hello "
const greeting_postfix = "!"
const french_prefix = "Bonjour "
const spanish_prefix = "Hola "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	var greet string
	switch lang {
	case "french":
		greet = french_prefix + name + greeting_postfix
	case "spanish":
		greet = spanish_prefix + name + greeting_postfix
	default:
		greet = english_prefix + name + greeting_postfix
	}
	return greet

}

func main() {
	fmt.Println(Hello("", ""))
}

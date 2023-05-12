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
	greeting := GreetingPrefix(lang) + name + greeting_postfix
	return greeting

}

func GreetingPrefix(lang string) string {
	var prefix string
	switch lang {
	case "french":
		prefix = french_prefix
	case "spanish":
		prefix = spanish_prefix
	default:
		prefix = english_prefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("", ""))
}

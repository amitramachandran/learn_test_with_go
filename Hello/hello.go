package main

import "fmt"

const (
	english_prefix   = "Hello "
	greeting_postfix = "!"
	french_prefix    = "Bonjour "
	spanish_prefix   = "Hola "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	greeting := greetingPrefix(lang) + name + greeting_postfix
	return greeting

}

func greetingPrefix(lang string) string {
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

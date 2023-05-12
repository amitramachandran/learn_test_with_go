package main

import "fmt"

const english_prefix = "Hello "
const greeting_postfix = "!"
const french_prefix = "Hola "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	var greet string
	if lang == "french" {
		greet = french_prefix + name + greeting_postfix
	} else {
		greet = english_prefix + name + greeting_postfix
	}
	return greet

}

func main() {
	fmt.Println(Hello("", ""))
}

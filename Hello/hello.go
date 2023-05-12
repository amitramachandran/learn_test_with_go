package main

import "fmt"

const greeting_prefix = "Hello "
const greeting_postfix = "!"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greeting_prefix + name + greeting_postfix

}

func main() {
	fmt.Println(Hello(""))
}

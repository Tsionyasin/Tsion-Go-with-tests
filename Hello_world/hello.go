package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHellowPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	if language == "French" {
		return frenchHellowPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}

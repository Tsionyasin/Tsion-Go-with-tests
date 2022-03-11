package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHellowPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHellowPrefix

	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix

	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}

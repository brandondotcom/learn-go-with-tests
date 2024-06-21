package main

import "fmt"

const (
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	spanishLanguageName = "Spanish"
	frenchHelloPrefix   = "Bonjour, "
	frenchLanguageName  = "French"
)

func Hello(name string, language string) string {

	if len(name) <= 0 {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanishLanguageName:
		prefix = spanishHelloPrefix
	case frenchLanguageName:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}

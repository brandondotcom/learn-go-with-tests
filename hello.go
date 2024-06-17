package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanishLanguageName = "Spanish"
const frenchHelloPrefix = "Bonjour, "
const frenchLanguageName = "French"

func Hello(name string, language string) string {

    if len(name) <= 0 { 
        name = "world" 
    }

    prefix := englishHelloPrefix

    switch language {
    case spanishLanguageName:
        prefix = spanishHelloPrefix
    case frenchLanguageName:
        prefix = frenchHelloPrefix
    }

    return prefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}


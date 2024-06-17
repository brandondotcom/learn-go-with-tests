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

    if language == spanishLanguageName {
        return spanishHelloPrefix + name
    }

    if language == frenchLanguageName {
        return frenchHelloPrefix + name
    }

    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}


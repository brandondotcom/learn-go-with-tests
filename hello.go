package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanishLanguageName = "Spanish"

func Hello(name string, language string) string {
    
    if len(name) <= 0 { 
        name = "world" 
    }

    if language == spanishLanguageName {
        return spanishHelloPrefix + name
    }

    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}


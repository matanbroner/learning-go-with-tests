package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const hebrew = "Hebrew"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const hebrewPrefix = "Shalom, "

func main() {
	fmt.Println(Hello("Matan", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case french:
		return frenchPrefix
	case spanish:
		return spanishPrefix
	case hebrew:
		return hebrewPrefix
	default:
		return englishPrefix
	}
}

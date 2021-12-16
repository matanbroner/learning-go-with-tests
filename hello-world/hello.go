package main

import "fmt"

func main() {
	fmt.Println(Hello("Matan"))
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	const englishPrefix = "Hello, "
	return englishPrefix + name
}

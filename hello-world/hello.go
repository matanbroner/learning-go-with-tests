package main

import "fmt"

func main() {
	fmt.Println(Hello("Matan"))
}

func Hello(name string) string {
	return "Hello, " + name
}

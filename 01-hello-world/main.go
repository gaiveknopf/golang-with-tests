package main

import "fmt"

const prefixHelloEnglish = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("%s%s!", prefixHelloEnglish, name)
}

func main() {
	fmt.Println(Hello("Charmaine"))
}

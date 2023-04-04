package main

import "fmt"

const (
	prefixHelloEnglish = "Hello, "
	prefixHelloSpanish = "Hola, "
	prefixHelloFrench  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		return fmt.Sprintf("%s%s!", prefixHelloSpanish, name)
	}
	if language == "French" {
		return fmt.Sprintf("%s%s!", prefixHelloFrench, name)
	}
	return fmt.Sprintf("%s%s!", prefixHelloEnglish, name)
}

func main() {
	fmt.Println(Hello("Charmaine", "Spanish"))
}

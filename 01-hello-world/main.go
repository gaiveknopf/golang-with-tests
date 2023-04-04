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
	return fmt.Sprintf("%s%s!", prefixGen(language), name)
}

func prefixGen(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = prefixHelloSpanish
	case "French":
		prefix = prefixHelloFrench
	default:
		prefix = prefixHelloEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("Charmaine", "Spanish"))
}

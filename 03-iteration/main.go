package main

import "fmt"

func Repeat(character string, quantity int) string {
	var repeated string
	for i := 0; i < quantity; i++ {
		repeated += character
	}
	return repeated
}

func main() {
	fmt.Println(Repeat("agora vai! ", 10))
}

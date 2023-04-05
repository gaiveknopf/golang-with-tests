package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("Hello, playground")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")

	result := make(chan int)
	go Add(1, 2, result)
	fmt.Println(<-result)

}

func Add(a, b int, result chan<- int) {
	result <- a + b
}

// A função Process executa uma operação e sinaliza a conclusão em um canal
var flag bool

func Process(done chan<- struct{}) {
	// Executa uma operação
	flag = true

	// Sinaliza a conclusão da operação no canal
	done <- struct{}{}
}

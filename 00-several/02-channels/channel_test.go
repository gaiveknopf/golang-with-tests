package main

import "testing"

func TestAdd(t *testing.T) {
	result := make(chan int)
	go Add(2, 3, result)
	if <-result != 5 {
		t.Errorf("expected 5 but got %d", <-result)
	}
}

func TestProcess(t *testing.T) {
	// Cria um canal para sincronizar a execução da função
	done := make(chan struct{})

	// Chama a função em uma goroutine, passando o canal como argumento
	go Process(done)

	// Espera pela sincronização do canal
	<-done

	// Verifica se a função foi executada corretamente
	if !flag {
		t.Errorf("A função Process não foi executada corretamente")
	}
}

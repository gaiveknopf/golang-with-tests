package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewMemoryPlayerStorage())

	if err := http.ListenAndServe(":5450", server); err != nil {
		log.Fatalf("não foi possível ouvir na porta 5000 %v", err)
	}
}

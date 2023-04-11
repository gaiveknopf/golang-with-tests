package main

import (
	api "golang-with-tests/15-api"
	"log"
	"net/http"
)

const dbFileName = "../../game.db.json"

func main() {
	storage, closer, err := api.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("fail on create storage %v", err)
	}
	defer closer()

	server := api.NewPlayerServer(storage)

	if err := http.ListenAndServe(":5450", server); err != nil {
		log.Fatalf("não foi possível ouvir na porta 5000 %v", err)
	}
}

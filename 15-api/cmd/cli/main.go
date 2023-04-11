package main

import (
	"fmt"
	api "golang-with-tests/15-api"
	"log"
	"os"
)

const dbFileName = "../../game.db.json"

func main() {

	storage, closer, err := api.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	defer closer()

	fmt.Println("Go Play Poker!")
	fmt.Println("Type {Name} wins to record a win")
	game := api.NewCLI(storage, os.Stdin)
	game.PlayPoker()
}

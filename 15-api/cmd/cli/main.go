package main

import (
	"fmt"
	api "golang-with-tests/15-api"
	"log"
	"os"
)

const dbFileName = "../../game.db.json"

func main() {
	store, _close, err := api.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer _close()

	game := api.NewTexasHoldem(api.BlindAlerterFunc(api.StdOutAlerter), store)
	cli := api.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("Let's play api")
	fmt.Println("Type {Name} wins to record a win")
	cli.PlayPoker()
}

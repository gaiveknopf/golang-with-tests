package api

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	playerStorage PlayerStore
	in            *bufio.Scanner
	out           io.Writer
	game          Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

const (
	PlayerPrompt         = "please enter the number of players: "
	BadPlayerInputErrMsg = "bad value received for number of players"
	BadWinnerInputErrMsg = "bad value received for winner"
)

func (cli *CLI) PlayPoker() {
	_, _ = fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayers, err := strconv.Atoi(cli.readLine())
	if err != nil {
		_, _ = fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}
	cli.game.Start(numberOfPlayers)
	winnerInput := cli.readLine()
	winner, err := extractWinner(winnerInput)
	if err != nil {
		_, _ = fmt.Fprint(cli.out, BadWinnerInputErrMsg)
		return
	}
	cli.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", fmt.Errorf(BadWinnerInputErrMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

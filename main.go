package main

import (
	"bufio"
	"console-games/gamesGo"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hi! This is Console Games!\nChoose the game you want to play!\n",
	" - Type 'h' if you want to play HANGMAN\n - Type 'p' if you want to play in PONG")
	inputGame := bufio.NewReader(os.Stdin)
	game, err := inputGame.ReadString('\n')
	if err != nil {
		panic(err)
	}
	game = strings.TrimSpace(game)
	switch game {
	case "h":
		gamesGo.HangmanFunc()
	case "p":
		gamesGo.Pong()
	}
	
}

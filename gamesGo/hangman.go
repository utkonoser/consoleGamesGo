package gamesGo

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)


func HangmanFunc() {
	rand.Seed(time.Now().Unix())
	words := []string{"golang", "hangman"}
	word := words[rand.Intn(len(words)-1)]
	var underscoreWord []string
	countOfMisses := 0
	for i := 0; i < len(word); i++ {
		underscoreWord = append(underscoreWord, " _")
	}
	fmt.Println(
		"THIS IS A HANGMAN GAME!\n",
		"You have to guess the word that I guess,\n",
		"otherwise, this will happen to you!\n",
		"  ____   \n",
		" |    |  \n",
		" |    O \n",
		" |   /|\\\n",
		" |   / \\ \n",
		"_|__death\n",
		"OK, let's play!")
	for countOfMisses < 6 {
		if strings.Join(underscoreWord, "") == word {
			fmt.Println(underscoreWord)
			fmt.Println("Congradulations! You escaped the hangman!")
			return
		}
		fmt.Printf("Word:%s\nLetter:", underscoreWord)
		inputLetter := bufio.NewReader(os.Stdin)
		letter, err := inputLetter.ReadString('\n')
		if err != nil {
			panic(err)
		}
		letter = strings.TrimSpace(letter)
		if strings.Contains(word, letter){
			for i, l := range word {
				if string(l) == letter {
					underscoreWord[i] = letter
				}
			}
		} else {
			countOfMisses++
			switch countOfMisses {
			case 1:
				fmt.Println(
					"  ____   \n",
					" |    |  \n",
					" |    O \n",
					" |   \n",
					" |     \n",
					"_|_____")
			case 2:
				fmt.Println(
					"  ____   \n",
					" |    |  \n",
					" |    O \n",
					" |    |\n",
					" |    \n",
					"_|______")
			case 3:
				fmt.Println(
					"  ____   \n",
					" |    |  \n",
					" |    O \n",
					" |   /|\n",
					" |    \n",
					"_|______")
			case 4:
				fmt.Println(
					"  ____   \n",
					" |    |  \n",
					" |    O \n",
					" |   /|\\\n",
					" |    \n",
					"_|______")
			case 5:
				fmt.Println(
					"  ____   \n",
					" |    |  \n",
					" |    O \n",
					" |   /|\\\n",
					" |   /  \n",
					"_|______")
			case 6:
				fmt.Println(
					"  ____   \n",
					" |    |  \n",
					" |    O \n",
					" |   /|\\\n",
					" |   / \\ \n",
					"_|__death\n",
					"You dead....")
				return
			}
		}
	}
}
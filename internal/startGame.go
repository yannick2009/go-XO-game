package internal

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-terminal-OX/pkg/constants"
	"github.com/go-terminal-OX/pkg/utils"
)

const clearSecond = 3
const pressEnterMsg = "Press ENTER to continue !"

var (
	ErrReadStdin = errors.New("error reading from stdin")
)

// InitGame prints the initial game message
func InitGame() {
	fmt.Print(constants.InitGameMsg)
	utils.ClearStdout(clearSecond)
	fmt.Print(constants.Rules)
	showRules()
}

// Show the game rules before to play
func showRules() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(pressEnterMsg + " ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(errors.Join(ErrReadStdin, err))
		}
		if input == "\n" {
			break
		}
	}
}

// ask to the player 1 to choose his symbol
func AskToChooseSymbol() {
	fmt.Println(utils.BorderText(constants.ChooseSymbolMsg))
}

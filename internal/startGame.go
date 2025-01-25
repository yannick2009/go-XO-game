package internal

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-OX-game/pkg/constants"
	"github.com/go-OX-game/pkg/utils"
	"github.com/go-OX-game/pkg/vars"
)

// InitGame prints the initial game message
func InitGame() {
	fmt.Print(constants.InitGameMsg)
	utils.ClearStdout(3)
	fmt.Print(vars.Rules)
	showRules()
}

// Show the game rules before to play
func showRules() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(constants.PressEnterMsg + " ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(errors.Join(errors.New("error reading from stdin"), err))
		}
		if input == "\n" {
			break
		}
	}
	utils.ClearStdout(0)
}

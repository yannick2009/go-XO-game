package internal

import (
	"fmt"

	"github.com/go-terminal-OX/pkg/constants"
	"github.com/go-terminal-OX/pkg/utils"
)

// Generate the delay output
func genDelayOutput(number string) {
	fmt.Println("The game start in: ")
	fmt.Print(number)
}

// DelayStart is a function that generates a delay before starting the game
func DelayStart() {
	delayNumbers := []string{constants.One, constants.Two, constants.Three}
	for _, num := range delayNumbers {
		genDelayOutput(num)
		utils.ClearStdout(1)
	}
	fmt.Print(constants.GoMsg)
	utils.ClearStdout(1)
}

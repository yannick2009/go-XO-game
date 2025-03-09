package internal

import (
	"fmt"

	"github.com/go-OX-game/pkg/constants"
	"github.com/go-OX-game/pkg/utils"
)

// Generate the delay output
func genDelayOutput(numStr string) {
	fmt.Println("The game start in: ")
	fmt.Print(numStr)
}

// DelayStart is a function that generates a delay before starting the game
func DelayStart() {
	utils.ClearStdout(1)
	delayNumbers := []string{constants.One, constants.Two, constants.Three}
	for _, numStr := range delayNumbers {
		genDelayOutput(numStr)
		utils.ClearStdout(1)
	}
	fmt.Print(constants.GoMsg)
}

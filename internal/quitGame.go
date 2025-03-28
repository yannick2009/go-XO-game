package internal

import (
	"fmt"
	"os"

	"github.com/go-OX-game/pkg/constants"
	"github.com/go-OX-game/pkg/utils"
)

// QuitGame is a function that quit the game
// and clear the screen
func QuitGame() {
	utils.ClearStdout(0)
	fmt.Println(constants.QuitGameMsg)
	os.Exit(0)
}

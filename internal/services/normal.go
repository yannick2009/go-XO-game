package services

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/go-OX-game/internal"
	"github.com/go-OX-game/pkg/constants"
	"github.com/go-OX-game/pkg/utils"
)

func LaunchNormalGame() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		utils.ClearStdout(0)
		fmt.Println(constants.QuitGameMsg)
		os.Exit(0)
	}()
	// Start the game
	internal.InitGame()
	// Show the game rules before to play
	internal.OkRules(nil)
	// Ask for the player one to choose his symbol between X or O
	internal.AskToChooseSymbol(nil)
	// a 3 seconds delay before starting the game
	internal.DelayStart(nil)
	// start the game
	internal.PlayGame()
}

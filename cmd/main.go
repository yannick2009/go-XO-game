package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/go-terminal-OX/internal"
	"github.com/go-terminal-OX/pkg/constants"
	"github.com/go-terminal-OX/pkg/utils"
)

func init() {
	// Clear the terminal
	utils.ClearStdout(0)
}

func main() {
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
	// Ask for the player one to choose his symbol between X or O
	internal.AskToChooseSymbol()
	// a 3 seconds delay before starting the game
	internal.DelayStart()
	// start the game
	internal.PlayGame()
}

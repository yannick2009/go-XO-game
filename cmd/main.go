package main

import (
	"github.com/go-terminal-OX/internal"
	"github.com/go-terminal-OX/pkg/utils"
)

func init() {
	// Clear the terminal
	utils.ClearStdout(0)
}

func main() {
	// Start the game
	internal.InitGame()
}

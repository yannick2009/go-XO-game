package internal

import (
	"fmt"

	"github.com/go-OX-game/pkg/constants"
	"github.com/go-OX-game/pkg/utils"
	"github.com/go-OX-game/pkg/vars"
)

// PlayGame is the function that start the game
func PlayGame() {
	for {
		var key string
		ShowGameTable()

		if CheckGameOver() {
			fmt.Println(utils.ArroundText("The game is over"))
			break
		}
		if CheckWinner() {
			if vars.ActivePlayer == 0 {
				fmt.Println(utils.BorderText("Player two (2) has won the game!"))
				break
			}
			fmt.Println(utils.BorderText("Player One (1) has won the game!"))
			break
		}

		for {
			activeplayerMsg()
			fmt.Scan(&key)
			if !IskeyValid(key) {
				continue
			}
			if !IsSquareAvailable(key) {
				fmt.Println("This square is already taken, choose another one")
				continue
			}
			break
		}

		Play(key)
		TurnPlayer()
		utils.ClearStdout(0)
	}
}

// play the game
func Play(key string) {
	val := vars.KeyToBoardMapping[key]
	if vars.ActivePlayer == 0 {
		*val = vars.PlayerOneSymbol
		return
	}
	*val = vars.PlayerTwoSymbol
}

// change the active player
func TurnPlayer() {
	vars.RemainingMoves--
	if vars.ActivePlayer == 0 {
		vars.ActivePlayer = 1
		return
	}
	vars.ActivePlayer = 0
}

// display the active player message
func activeplayerMsg() {
	if vars.ActivePlayer == 0 {
		fmt.Println("=> PLAYER 1")
		fmt.Println(constants.ValidkeysMsg)
		return
	}
	fmt.Println("=> PLAYER 2")
	fmt.Println(constants.ValidkeysMsg)
}

// check if is key entered is valid
func IskeyValid(key string) bool {
	for _, v := range vars.ValidKeysArr {
		if key == v {
			return true
		}
	}
	return false
}

// check if the game table square is available
func IsSquareAvailable(key string) bool {
	return *vars.KeyToBoardMapping[key] == ""
}

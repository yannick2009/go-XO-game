package internal

import (
	"fmt"

	"github.com/go-OX-game/pkg/config"
	"github.com/go-OX-game/pkg/constants"
	"github.com/go-OX-game/pkg/types"
	"github.com/go-OX-game/pkg/utils"
	"github.com/go-OX-game/pkg/vars"
	"slices"
)

// PlayGame is the function that start the game
// it is used for the normal game
func PlayGame() {
	utils.ClearStdout(1)
	for {
		var key string
		ShowGameTable()

		// check if the game is over or if there is a winner
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

		// make the current player choose a square
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

		// play the game with the given key and symbol
		if vars.ActivePlayer == 0 {
			Play(key, vars.PlayerOneSymbol)
		} else {
			Play(key, vars.PlayerTwoSymbol)
		}
		// turn to the next player
		TurnPlayer()
		utils.ClearStdout(0)
	}
}

// PlayGameWS is the function that start the game with websocket
// it is used for the client and server
func PlayGameWS(ws *config.WS, isMyTurn bool, playerSymbol string) {
	utils.ClearStdout(0)

	// check if the game is over or if there is a winner
	if CheckGameOver() {
		fmt.Println(utils.ArroundText("The game is over"))
		return
	}
	if CheckWinner() {
		if vars.MyTurn {
			fmt.Println(utils.BorderText("Player two (2) has won the game!"))
			return
		}
		fmt.Println(utils.BorderText("You have won the game!"))
		return
	}

	var key string
	ShowGameTable()

	// check if it is my turn and play the game
	// if it is not my turn, wait for the other player to play
	if isMyTurn {
		for {
			fmt.Println("Your turn")
			fmt.Println(constants.ValidkeysMsg)

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


		Play(key, playerSymbol)
		msg, err := utils.EncodeGameDataMsg(types.GameDataMessage{
			IsMyTurn:  true,
			GameBoard: vars.GameBoard,
		})
		if err != nil {
			fmt.Println("Error when encoding message")
			return
		}
		vars.MyTurn = false
		ws.SendMsg(types.Message{
			Type: constants.PLAY,
			Data: msg,
		})

		// check if the game is over or if there is a winner
		utils.ClearStdout(0)
		if CheckGameOver() {
			fmt.Println(utils.ArroundText("The game is over"))
			return
		} else if CheckWinner() {
			if vars.MyTurn {
				fmt.Println(utils.BorderText("Player two (2) has won the game!"))
				return
			}
			fmt.Println(utils.BorderText("You have won the game!"))
			return
		} else {
			ShowGameTable()
		}
	}
}

// play the game with the given key and symbol
func Play(key, symbol string) {
	val := vars.KeyToBoardMapping[key]
	*val = symbol
}

// change the active player
func TurnPlayer() {
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
	return slices.Contains(vars.ValidKeysArr, key)
}

// check if the game table square is available
func IsSquareAvailable(key string) bool {
	return *vars.KeyToBoardMapping[key] == ""
}

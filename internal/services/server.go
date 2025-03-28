package services

import (
	"log"
	"os"
	"os/signal"

	"github.com/go-OX-game/internal"
	"github.com/go-OX-game/pkg/config"
	"github.com/go-OX-game/pkg/constants"
	"github.com/go-OX-game/pkg/types"
	"github.com/go-OX-game/pkg/utils"
	"github.com/go-OX-game/pkg/vars"
)

const ErrorDecodingClientMsg = "Error when decoding server message received"

func LunchServerGame(ws *config.WS, channel *chan types.MsgChan) {
	vars.MyTurn = true
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		ws.SendMsg(types.Message{
			Type: constants.QUIT,
			Data: nil,
		})
		internal.QuitGame()
	}()

	// Start the game
	internal.InitGame()
	// Validate the game rules
	internal.OkRules(ws)
	// Ask for the player one to choose his symbol between X or O
	internal.AskToChooseSymbol(ws)
	// a 3 seconds delay before starting the game
	internal.DelayStart(ws)
	// Start the game
	internal.PlayGameWS(ws, vars.MyTurn, vars.PlayerOneSymbol)

	for msg := range *channel {
		switch msg.Type {
		case constants.PLAY:
			data, err := utils.DecodeGameDataMsg(msg.Msg)
			if err != nil {
				log.Print(ErrorDecodingClientMsg)
			}
			PopulateGame(data)
			internal.PlayGameWS(ws, vars.MyTurn, vars.PlayerOneSymbol)
		case constants.QUIT:
			internal.QuitGame()
		}
	}
}

// populateGameBoard is a function that populate the game board when a ws message is received
func PopulateGame(data types.GameDataMessage) {
	vars.MyTurn = data.IsMyTurn
	vars.GameBoard = data.GameBoard
}

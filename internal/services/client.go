package services

import (
	"fmt"
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

const ErrorDecodingServerMsg = "Error when decoding server message received"

func LunchClientGame(ws *config.WS, channel *chan types.MsgChan) {
	vars.MyTurn = false
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Listen for interrupt signal to quit the game
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

	// listen to the channel for messages from the server
	// and handle them accordingly
	for msg := range *channel {
		switch msg.Type {
		// case when the game rules are readed
		case constants.RULES_READED:
			utils.ClearStdout(0)
			fmt.Println("Waiting for the other player to choose his symbol...")
		// case when players are asked to choose their symbol
		case constants.CHOOSE_SYMBOL:
			utils.ClearStdout(0)
			fmt.Println("Your symbol is: " + string(msg.Msg))
			vars.PlayerTwoSymbol = string(msg.Msg)
		// case when the game starts
		case constants.START_GAME:
			internal.DelayStart(nil)
			internal.PlayGameWS(ws, vars.MyTurn, vars.PlayerTwoSymbol)
		// case when the game is played
		case constants.PLAY:
			data, err := utils.DecodeGameDataMsg(msg.Msg)
			if err != nil {
				log.Print(ErrorDecodingServerMsg)
			}
			PopulateGame(data)
			internal.PlayGameWS(ws, data.IsMyTurn, vars.PlayerTwoSymbol)
		// case when the game is over by the server who quits the game
		case constants.QUIT:
			internal.QuitGame()
		}
	}
}

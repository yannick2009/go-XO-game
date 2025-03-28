package handlers

import (
	"encoding/json"
	"log"

	"github.com/go-OX-game/internal/services"
	"github.com/go-OX-game/pkg/config"
	"github.com/go-OX-game/pkg/types"
)

// ClientHandler is a function that handle the client websocket connection
func ClientHandler(ws *config.WS) {
	defer ws.Conn.Close()

	channel := make(chan types.MsgChan) // Channel to receive messages from the server

	// Start a goroutine to receive messages from the server
	go func() {
		for {
			_, msgRaw, err := ws.ReceiveMsg()
			if err != nil {
				log.Println("Error on receiving ws message")
				break
			}

			msg := &types.Message{}
			err = json.Unmarshal(msgRaw, msg)
			if err != nil {
				log.Println("Error on unmarshalling ws message")
				break
			}

			channel <- types.MsgChan{
				Type: msg.Type,
				Msg:  msg.Data,
			}
		}
	}()

	// Start the client game
	services.LunchClientGame(ws, &channel)
}

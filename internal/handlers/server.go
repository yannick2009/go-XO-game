package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-OX-game/internal/services"
	"github.com/go-OX-game/pkg/config"
	"github.com/go-OX-game/pkg/types"
)

// ServerHandler is a function that handle the server websocket connection
func ServerHandler(w http.ResponseWriter, r *http.Request) {
	ws := config.NewServerWS(w, r) // Create a new websocket instance
	defer ws.Conn.Close()

	channel := make(chan types.MsgChan) // Channel to receive messages from the client

	// Start a goroutine to receive messages from the client
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

	// Start the server game
	services.LunchServerGame(ws, &channel)
}

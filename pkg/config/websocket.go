package config

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-OX-game/pkg/types"
	"github.com/gorilla/websocket"
)

var (
	ErrCreateWSInstance = errors.New("creation of websocket instance failed")
)

type WS struct {
	Conn *websocket.Conn
}

// create a websocket handler for server
func NewServerWS(w http.ResponseWriter, r *http.Request) *WS {
	// Upgrader is a websocket upgrader
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(errors.Join(ErrCreateWSInstance, err))
	}
	return &WS{
		Conn: c,
	}
}

// create a websocket handler for client
func NewClientWS(conn *websocket.Conn) *WS {
	return &WS{
		Conn: conn,
	}
}

// send message to client
func (w *WS) SendMsg(msg types.Message) error {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("Error when encoding message to JSON")
	}

	return w.Conn.WriteMessage(websocket.BinaryMessage, jsonMsg)
}

// receive message to client
func (w *WS) ReceiveMsg() (int, []byte, error) {
	return w.Conn.ReadMessage()
}

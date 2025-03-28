package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-OX-game/internal/handlers"
	"github.com/go-OX-game/internal/services"
	"github.com/go-OX-game/pkg/config"
	"github.com/go-OX-game/pkg/utils"
	"github.com/gorilla/websocket"
)

func init() {
	// Clear the terminal
	utils.ClearStdout(0)
}

func main() {
	// Parse command line arguments for client
	serverURL := flag.String("url", "", "Server URL")
	client := flag.Bool("client", false, "Run the client")
	// Parse command line arguments for server
	server := flag.Bool("server", false, "Run the server")
	flag.Parse()

	if *server {
		http.HandleFunc("/ws", handlers.ServerHandler)
		log.Printf("Serveur WebSocket démarré sur ws://%s:3000/ws", utils.GetMyIP())
		log.Fatal(http.ListenAndServe(":3000", nil))
		return
	} else if *client {
		conn, _, err := websocket.DefaultDialer.Dial(*serverURL, nil)
		if err != nil {
			log.Fatal("Error when connecting to the server")
		}
		ws := config.NewClientWS(conn)
		handlers.ClientHandler(ws)
		return
	} else {
		services.LaunchNormalGame()
	}

}

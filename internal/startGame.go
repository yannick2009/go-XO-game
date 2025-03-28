package internal

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-OX-game/pkg/config"
	"github.com/go-OX-game/pkg/constants"
	"github.com/go-OX-game/pkg/types"
	"github.com/go-OX-game/pkg/utils"
	"github.com/go-OX-game/pkg/vars"
)

// InitGame prints the initial game message
func InitGame() {
	utils.ClearStdout(0)
	fmt.Print(constants.InitGameMsg)
	utils.ClearStdout(3)
	fmt.Print(vars.Rules)
}

// OkRules is a function that wait for the user to press enter
// to continue the game and send a message to the server
func OkRules(ws *config.WS) {
	fmt.Printf(constants.PressEnterMsg + " ")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(errors.Join(errors.New("error reading from stdin"), err))
		}
		if input == "\n" {
			if ws != nil {
				ws.SendMsg(types.Message{
					Type: constants.RULES_READED,
					Data: nil,
				})
			}
			break
		}
	}
	utils.ClearStdout(0)
}

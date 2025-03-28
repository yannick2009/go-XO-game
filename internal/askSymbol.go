package internal

import (
	"fmt"
	"strings"

	"github.com/go-OX-game/pkg/config"
	"github.com/go-OX-game/pkg/constants"
	"github.com/go-OX-game/pkg/types"
	"github.com/go-OX-game/pkg/utils"
	"github.com/go-OX-game/pkg/vars"
)

// ask to the player 1 to choose his symbol
func AskToChooseSymbol(ws *config.WS) {
	utils.ClearStdout(0)
	// Define the symbols
	xSymbol := "x"
	oSymbol := "o"

	// Ask to the player one to choose his symbol
	var symbolChoosed string
	fmt.Println(utils.BorderText(constants.ChooseSymbolMsg))
	fmt.Scan(&symbolChoosed)

	// Check if the player one choose a valid symbol
	for strings.ToLower(symbolChoosed) != xSymbol && strings.ToLower(symbolChoosed) != oSymbol {
		fmt.Println("You can only choose between X or O !")
		fmt.Scan(&symbolChoosed)
	}

	// Show the symbol choosed by the player one
	phraseTmpl := "Player one is %s and player two is %s\n"
	switch strings.ToLower(symbolChoosed) {
	case xSymbol:
		fmt.Printf(phraseTmpl, strings.ToUpper(xSymbol), strings.ToUpper(oSymbol))
		vars.PlayerOneSymbol = xSymbol
		vars.PlayerTwoSymbol = oSymbol
	case oSymbol:
		fmt.Printf(phraseTmpl, strings.ToUpper(oSymbol), strings.ToUpper(xSymbol))
		vars.PlayerOneSymbol = oSymbol
		vars.PlayerTwoSymbol = xSymbol
	}
	if ws != nil {
		ws.SendMsg(types.Message{
			Type: constants.CHOOSE_SYMBOL,
			Data: []byte(vars.PlayerTwoSymbol),
		})
	}
}

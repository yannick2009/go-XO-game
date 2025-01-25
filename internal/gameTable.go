package internal

import (
	"fmt"
	"strings"

	"github.com/go-terminal-OX/pkg/vars"
)

// GameTable is the game table
var rowTop = strings.Repeat("┌───┐", 3) + "\n"
var rowbottom = strings.Repeat("└───┘", 3) + "\n"

// genRowMiddle generate the middle row
func genRowMiddle(rowNum int) string {
	result := ""
	for i := range 3 {
		if len(vars.GameBoard[rowNum][i]) == 0 {
			result += fmt.Sprintf("│ %s │", " ")
			continue
		}
		result += fmt.Sprintf("│ %s │", vars.GameBoard[rowNum][i])
	}
	return result + "\n"
}

// GameTable is the game table
func generateGameTable() string {
	var result string
	for i := range 3 {
		result += rowTop + genRowMiddle(i) + rowbottom
	}
	return result
}

// ShowGameTable show the game table
func ShowGameTable() {
	fmt.Print(generateGameTable())
}

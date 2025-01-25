package vars

import _ "embed"

var (
	//go:embed rules.txt
	Rules           string     // Rules of the game
	RemainingMoves  uint   = 9 // Party action
	ActivePlayer    int        // Active player
	PlayerOneSymbol string     // Player one symbol
	PlayerTwoSymbol string     // Player two symbol

	ValidKeysArr      []string           = []string{"a", "z", "e", "q", "s", "d", "w", "x", "c"} // Valid keys list
	GameBoard         [3][3]string                                                               // Game board
	KeyToBoardMapping map[string]*string = map[string]*string{                                   // Key to board mapping
		"a": &GameBoard[0][0],
		"z": &GameBoard[0][1],
		"e": &GameBoard[0][2],
		"q": &GameBoard[1][0],
		"s": &GameBoard[1][1],
		"d": &GameBoard[1][2],
		"w": &GameBoard[2][0],
		"x": &GameBoard[2][1],
		"c": &GameBoard[2][2],
	}
)

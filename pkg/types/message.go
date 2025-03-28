package types

// Package types contains the types used in the game
type Message struct {
	Type int
	Data []byte
}

// GameDataMessage is a struct that contains the game data
type GameDataMessage struct {
	IsMyTurn  bool
	GameBoard [3][3]string
}

// MsgChan is a struct that contains the type of the message and the message itself
type MsgChan struct {
	Type int
	Msg  []byte
}

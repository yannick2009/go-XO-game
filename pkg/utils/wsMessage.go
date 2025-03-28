package utils

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/go-OX-game/pkg/types"
)

var (
	ErrEncode = errors.New("message encoding failed")
	ErrDecode = errors.New("message decoding failed")
)

// encode ws message (struct) in gob then in byte
func EncodeGameDataMsg(msg types.GameDataMessage) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(msg)
	if err != nil {
		return nil, errors.Join(ErrEncode, err)
	}

	return buf.Bytes(), nil
}

// decode ws message ([]byte) received in gob then in struct
func DecodeGameDataMsg(msg []byte) (types.GameDataMessage, error) {
	var message types.GameDataMessage
	buf := bytes.NewBuffer(msg)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(&message)
	if err != nil {
		return types.GameDataMessage{}, errors.Join(ErrDecode, err)
	}

	return message, nil
}

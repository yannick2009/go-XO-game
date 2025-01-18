package utils

import (
	"fmt"
	"time"
)

// ClearStdout clears the terminal after t seconds
func ClearStdout(t uint) {
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Print("\033[H\033[2J")
}

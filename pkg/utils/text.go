package utils

import "strings"

// BorderText returns a string with a border around the text
func BorderText(text string) string {
	paddingWidth := 4
	borderChar := "-"
	horizontalPadding := strings.Repeat(borderChar, paddingWidth)
	verticalPadding := strings.Repeat(" ", paddingWidth)
	horizontalBorder := "+" + horizontalPadding + strings.Repeat(borderChar, len(text)) + horizontalPadding + "+\n"

	return horizontalBorder + "|" + verticalPadding + text + verticalPadding + "|\n" + horizontalBorder
}

// ArroundText returns a string with a border around the text
func ArroundText(text string) string {
	text = strings.ToUpper(text)
	borderSymbol := "="
	symbolCount := 5
	border := strings.Repeat(borderSymbol, symbolCount)
	
	return border + " " + text + " " + border
}

package constants

const InitGameMsg = `
__  _____         ____    _    __  __ _____ 
\ \/ / _ \       / ___|  / \  |  \/  | ____|
 \  / | | |_____| |  _  / _ \ | |\/| |  _|  
 /  \ |_| |_____| |_| |/ ___ \| |  | | |___ 
/_/\_\___/       \____/_/   \_\_|  |_|_____|

		Welcome !
We are going to show you the rule in few seconds.


    Developed with ❤️ by: @yannick2009
` // Initial game message

const (
	QuitGameMsg     = "Thanks for playing with us! Goodbye!" // Quit game message`
	ChooseSymbolMsg = "Choose your symbol: X or O"           // Ask the player one to choose his symbol message
	ValidkeysMsg    = "Valid Keys: a,z,e,q,s,d,w,x,c"        // Valid keys message
	PressEnterMsg   = "Press ENTER to continue !"            // Press enter to continue message

	One   = " _ \n/ |\n| |\n| |\n|_|"                                                                // One in ascii art
	Two   = " ____  \n|___ \\ \n  __) |\n / __/ \n|_____|"                                           // Two in ascii art
	Three = " _____ \n|___ / \n  |_ \\ \n ___) |\n|____/ "                                           // Three in ascii art
	GoMsg = "  ____  ___  _ \n / ___|/ _ \\| |\n| |  _| | | | |\n| |_| | |_| |_|\n \\____|\\___|(_)" // Go! in ascii art
)

// ws message types
const (
	RULES_READED = iota
	CHOOSE_SYMBOL
	START_GAME
	PLAY
	QUIT
	// END
)

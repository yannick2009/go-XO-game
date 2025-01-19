package constants

import _ "embed"

const InitGameMsg = `
__  _____         ____    _    __  __ _____ 
\ \/ / _ \       / ___|  / \  |  \/  | ____|
 \  / | | |_____| |  _  / _ \ | |\/| |  _|  
 /  \ |_| |_____| |_| |/ ___ \| |  | | |___ 
/_/\_\___/       \____/_/   \_\_|  |_|_____|

		Welcome !
We are going to show you the rule in few seconds.
`

const ChooseSymbolMsg = "Choose your symbol: X or O"

//go:embed rules.txt
var Rules string

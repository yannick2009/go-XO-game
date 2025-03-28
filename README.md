```
__  _____         ____    _    __  __ _____ 
\ \/ / _ \       / ___|  / \  |  \/  | ____|
 \  / | | |_____| |  _  / _ \ | |\/| |  _|  
 /  \ |_| |_____| |_| |/ ___ \| |  | | |___ 
/_/\_\___/       \____/_/   \_\_|  |_|_____|
```

# TIC TAC TOE - Terminal Game

A Go implementation of Tic Tac Toe that can be played in the terminal with multiple game modes.

## DESCRIPTION

This project is a Tic Tac Toe game developed entirely in Golang that can be played directly in the terminal. The game offers three different play modes:

- **Normal mode**: Two players compete on the same machine
- **Server mode**: Hosts a game that a client can connect to
- **Client mode**: Allows joining a game hosted on a server

The game is played on a 3x3 grid where players take turns placing their mark (X or O). The first player to align three of their marks (horizontally, vertically, or diagonally) wins the game. If the grid is filled without a winner, the game is declared a draw.

## INSTALLATION

### Prerequisites

- Go (version 1.16 or higher): [https://golang.org/dl/](https://golang.org/dl/)

### Download

Clone the repository to your machine:

```bash
git clone https://github.com/yannick2009/go-terminal-OX.git
cd go-terminal-OX
```

### Compilation

Compile the game with the following command:

```bash
go build -o xo-game cmd/main.go 
```

## USAGE

### Normal mode (two local players)

To start a local game with two players:

```bash
./xo-game
```

### Server mode

#### Prererequisites

You must be in the same network as the client. Make sure to allow WebSocket connections on your firewall.

To start the game in server mode, send the ip address printed to the client (2nd player) and wait for his connection:

```bash
./xo-game --server
```

### Client mode

To join a game hosted on a server:

```bash
./xo-game -client --url=ws://<SERVER_IP_ADDRESS>/ws
```

## CONTRIBUTING

Contributions are welcome! If you would like to contribute to this project:

1. Fork the repository
2. Create a branch for your feature (`git checkout -b new-feature`)
3. Commit your changes (`git commit -m 'Add new feature'`)
4. Push to the branch (`git push origin new-feature`)
5. Create a Pull Request

If you find bugs or have suggestions for improvements, feel free to open an issue.

## MAINTAINERS

This project is maintained by: Yannick Kouakou 
[LinkedIn](https://www.linkedin.com/in/yannick-k-946970200/) | [GitHub](https://github.com/yannick2009/)

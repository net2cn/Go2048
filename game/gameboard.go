package game

import (
	"fmt"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type GameBoard struct {
	lastBoard [4][4]int
	board     [4][4]int

	GameOverFlag bool
}

func NewGameBoard() *GameBoard {
	fmt.Println("Creating new board.")
	gameBoard := GameBoard{}
	gameBoard.GameOverFlag = false
	gameBoard.generateNewTile()
	gameBoard.PrintBoard()
	return &gameBoard
}

func (gameBoard *GameBoard) PrintBoard() {
	for i := 0; i < 4; i++ {
		fmt.Println(gameBoard.board[i])
	}
}

func (gameBoard *GameBoard) generateNewTile() {
	var emptyTiles []([2]int)

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if gameBoard.board[x][y] == 0 {
				cord := [2]int{x, y}
				emptyTiles = append(emptyTiles, cord)
			}
		}
	}

	if len(emptyTiles) < 1 {
		gameBoard.GameOverFlag = true
		return
	}

	maximum := rand.Intn(2 + 1)
	for i := 0; i < maximum; i++ {
		idx := rand.Intn(len(emptyTiles))
		value := (rand.Intn(2) + 1) * 2

		cord := emptyTiles[idx]
		gameBoard.board[cord[0]][cord[1]] = value
	}
}

func (gameBoard *GameBoard) rotateBoard(times int) {
	for i := 0; i < 4; i++ {
		copy(gameBoard.lastBoard[i][0:], gameBoard.board[i][0:])
	}

	for i := 0; i < times; i++ {
		for x := 0; x < 4/2; x++ {
			for y := x; y < 4-x-1; y++ {
				temp := gameBoard.board[x][y]
				gameBoard.board[x][y] = gameBoard.board[4-1-y][x]
				gameBoard.board[4-1-y][x] = gameBoard.board[4-1-x][4-1-y]
				gameBoard.board[4-1-x][4-1-y] = gameBoard.board[y][4-1-x]
				gameBoard.board[y][4-1-x] = temp
			}
		}
	}
}

func (gameBoard *GameBoard) Update(keyState sdl.Keycode) {
	input := -1

	switch keyState {
	case sdl.K_UP:
		fmt.Println("Up pressed.")
		input = 1
	case sdl.K_DOWN:
		fmt.Println("Down pressed.")
		input = 3
	case sdl.K_LEFT:
		fmt.Println("Left pressed.")
		input = 0
	case sdl.K_RIGHT:
		fmt.Println("Rigth pressed.")
		input = 2
	}

	if input != -1 && gameBoard.GameOverFlag != true {
		gameBoard.rotateBoard(input)
		gameBoard.generateNewTile()
		gameBoard.PrintBoard()
	}

	if gameBoard.GameOverFlag == true {
		fmt.Println("Game over!")
	}
}

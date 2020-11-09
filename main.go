package main

import (
	"fmt"

	"github.com/net2cn/Go2048/game"
)

// SDL2 variables
var windowTitle string = "Go2048 SDL2"
var fontPath string = "./assets/UbuntuMono-R.ttf" // Man, do not use a variable-width font! It looks too ugly with that!
var fontSize int = 20
var windowWidth, windowHeight int32 = 640, 480

func main() {
	fmt.Println(windowTitle)
	fmt.Println("Yet another 2048 game written in golang.")

	game := game.NewController(windowWidth, windowHeight, fontPath, fontSize, windowTitle)
	game.Start()

	fmt.Println("Game exited. Bye!")
}

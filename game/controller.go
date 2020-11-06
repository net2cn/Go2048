package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

var gameTitle string

type Controller struct {
	Renderer  *Renderer
	GameBoard *GameBoard

	exitFlag  bool
	inputLock bool
}

func NewController(width int32, height int32, fontPath string, fontSize int, windowTitle string) *Controller {
	controller := Controller{
		Renderer:  NewRenderer(width, height, fontPath, fontSize, windowTitle),
		GameBoard: NewGameBoard(),

		exitFlag: false,
	}

	gameTitle = windowTitle

	return &controller
}

func (controller *Controller) Update() {
	keyState := sdl.GetKeyboardState()
	if keyState[sdl.SCANCODE_ESCAPE] != 0 {
		controller.exitFlag = true
		return
	}

	// Get inputs.
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			controller.exitFlag = true
			return
		case *sdl.KeyboardEvent:
			if !controller.inputLock {
				controller.GameBoard.Update(t.Keysym.Sym)
			}

			// Anti-jittering
			if t.Repeat > 0 {
				// Held.
				controller.inputLock = false
			} else {
				// Pressed once.
				if t.State == sdl.RELEASED {
					controller.inputLock = false
				} else if t.State == sdl.PRESSED {
					controller.inputLock = true
				}
			}
		}
	}

	controller.Renderer.Update(*controller.GameBoard)

	controller.exitFlag = false
}

func (controller *Controller) Start() {
	for !controller.exitFlag {
		controller.Update()
	}
}
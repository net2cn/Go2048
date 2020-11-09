package game

import (
	"fmt"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var colorTable map[string]uint32 = map[string]uint32{
	"Background": 0x00fbf8ef,
	"GameBoard":  0x00bbada0,
	"Tile0":      0xc4eee4da,
	"Tile2":      0x00eee4da,
	"Tile4":      0x00ede0c8,
	"Tile8":      0x00f2b179,
	"Tile16":     0x00f59563,
	"Tile32":     0x00f67c5f,
	"Tile64":     0x00f65e3b,
	"Tile128":    0x00edcf72,
	"Tile256":    0x00edcc61,
	"Tile512":    0x00edc850,
	"Tile1024":   0x00edc53f,
	"Tile2048":   0x00edc22e,
}

type Renderer struct {
	window   *sdl.Window
	surface  *sdl.Surface
	buffer   *sdl.Surface
	renderer *sdl.Renderer
	font     *ttf.Font
}

func NewRenderer(width int32, height int32, fontPath string, fontSize int, windowTitle string) *Renderer {
	var err error

	renderer := Renderer{
		window:   nil,
		surface:  nil,
		buffer:   nil,
		renderer: nil,
		font:     nil,
	}

	// Initialize sdl2
	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Printf("Failed to init sdl2: %s\n", err)
		panic(err)
	}

	// Initialize font
	if err = ttf.Init(); err != nil {
		fmt.Printf("Failed to init font: %s\n", err)
		panic(err)
	}

	// Load the font for our text
	if renderer.font, err = ttf.OpenFont(fontPath, fontSize); err != nil {
		fmt.Printf("Failed to load font: %s\n", err)
		panic(err)
	}

	// Create window
	renderer.window, err = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Failed to create window: %s\n", err)
		panic(err)
	}

	// Create draw surface and draw buffer
	if renderer.surface, err = renderer.window.GetSurface(); err != nil {
		fmt.Printf("Failed to get window surface: %s\n", err)
		panic(err)
	}

	if renderer.buffer, err = renderer.surface.Convert(renderer.surface.Format, renderer.window.GetFlags()); err != nil {
		fmt.Printf("Failed to create buffer: %s\n", err)
	}

	// Create renderer
	renderer.renderer, err = sdl.CreateRenderer(renderer.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("Failed to create renderer: %s\n", err)
		panic(err)
	}

	return &renderer
}

// Private draw functions
func (renderer *Renderer) drawString(x int, y int, str string, color *sdl.Color) {
	if len(str) <= 0 {
		return
	}

	// Create text
	text, err := renderer.font.RenderUTF8Blended(str, *color)
	if err != nil {
		fmt.Printf("Failed to create text: %s\n", err)
		return
	}
	defer text.Free()

	// Draw text, noted that you should always draw on buffer instead of directly draw on screen and blow yourself up.
	if err = text.Blit(nil, renderer.buffer, &sdl.Rect{X: int32(x), Y: int32(y)}); err != nil {
		fmt.Printf("Failed to draw text: %s\n", err)
		return
	}
}

// Draw a full sprite.
func (renderer *Renderer) drawSprite(x int, y int, sprite *sdl.Surface) {
	sprite.Blit(nil, renderer.buffer, &sdl.Rect{X: int32(x), Y: int32(y)})
}

// Draw a part of sprite.
func (renderer *Renderer) drawPartialSprite(dstX int, dstY int, sprite *sdl.Surface, srcX int, srcY int, w int, h int) {
	dstRect := sdl.Rect{X: int32(dstX), Y: int32(dstY), W: int32(w), H: int32(h)}
	srcRect := sdl.Rect{X: int32(srcX), Y: int32(srcY), W: int32(w), H: int32(h)}
	sprite.Blit(&srcRect, renderer.buffer, &dstRect)
}

func (renderer *Renderer) Update(gameBoard GameBoard) {
	// Render game.
	// 0xAARRGGBB
	renderer.buffer.FillRect(&sdl.Rect{X: 0, Y: 0, W: renderer.buffer.W, H: renderer.buffer.H}, colorTable["Background"])                        // Fill Background.
	renderer.buffer.FillRect(&sdl.Rect{X: 0 + 15, Y: 0 + 15, W: renderer.buffer.H - 15*2, H: renderer.buffer.H - 15*2}, colorTable["GameBoard"]) // Fill Game Board Background.
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			// Render tiles.
			tile := strconv.Itoa(gameBoard.board[y][x])
			tileName := "Tile" + tile
			renderer.buffer.FillRect(&sdl.Rect{X: int32(0 + 15 + 10 + x*110), Y: int32(0 + 15 + 10 + y*110), W: 100, H: 100}, colorTable[tileName])
			if tile != "0" {
				renderer.drawString(0+15+10+x*110+45-((len(tile)-1)*5), 0+15+10+y*110+45, tile, &sdl.Color{R: 0, G: 0, B: 0, A: 0})
			}
		}
	}
	renderer.drawString(480, 15, "Score", &sdl.Color{R: 0, G: 0, B: 0, A: 0})
	renderer.drawString(480, 30, strconv.Itoa(gameBoard.GameScore), &sdl.Color{R: 0, G: 0, B: 0, A: 0})

	if gameBoard.GameOverFlag == true {
		renderer.drawString(320-25, 240-10, "Game Over!", &sdl.Color{R: 0, G: 0, B: 0, A: 0})
		renderer.drawString(320-75, 240+10, "Press \"R\" to restart.", &sdl.Color{R: 0, G: 0, B: 0, A: 0})
	}

	if gameBoard.AccomplishedFlag == true && gameBoard.ContinueFlag == false {
		renderer.drawString(320-20, 240-10, "You Won!", &sdl.Color{R: 0, G: 0, B: 0, A: 0})
		renderer.drawString(320-90, 240+10, "Press \"Enter\" to continue.", &sdl.Color{R: 0, G: 0, B: 0, A: 0})
	}

	// Swap buffer and present our rendered content.
	renderer.window.UpdateSurface()
	renderer.buffer.Blit(nil, renderer.surface, nil)

	// Clear out buffer for next render round.
	renderer.buffer.FillRect(nil, 0xFF000000)
	renderer.renderer.Clear()
}

package game

import "github.com/veandco/go-sdl2/sdl"

const (
	screenWidth  = 600
	screenHeight = 600
)

var Window *sdl.Window

func InitWindow() error {
	var err error
	Window, err = sdl.CreateWindow("Go Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	return err
}

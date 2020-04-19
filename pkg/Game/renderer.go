package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var Renderer *sdl.Renderer

func InitRenderer() error {
	var err error
	Renderer, err = sdl.CreateRenderer(Window, 1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		return fmt.Errorf("creating renderer %v", err)
	}
	return nil
}

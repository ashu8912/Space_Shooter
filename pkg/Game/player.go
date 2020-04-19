package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	Tex         *sdl.Texture
	ImageWidth  int32
	ImageHeight int32
}

func NewPlayer() (Player, error) {
	var p Player
	surface, err := sdl.LoadBMP("sprites/player.bmp")
	defer surface.Free()
	if err != nil {
		return Player{}, fmt.Errorf("loading image %v", err)
	}
	imageWidth := surface.W
	imageHeight := surface.H
	playerTexture, err := Renderer.CreateTextureFromSurface(surface)

	if err != nil {
		return Player{}, fmt.Errorf("creating texture %v", err)
	}
	p.ImageHeight = imageHeight
	p.ImageWidth = imageWidth
	p.Tex = playerTexture
	return p, nil
}

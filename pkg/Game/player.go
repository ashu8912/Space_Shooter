package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed float64 = 0.1
)

type Player struct {
	Tex         *sdl.Texture
	imageWidth  int32
	imageHeight int32
	x, y        float64
}

func (p *Player) Draw() {
	Renderer.Copy(p.Tex, &sdl.Rect{
		X: 0,
		Y: 0,
		W: p.imageWidth,
		H: p.imageHeight,
	}, &sdl.Rect{
		X: (int32)(p.x),
		Y: (int32)(p.y),
		W: p.imageWidth,
		H: p.imageHeight,
	})
}
func (p *Player) UpdatePlayerPos() {
	keys := sdl.GetKeyboardState()
	var x float64
	if keys[sdl.SCANCODE_LEFT] == 1 {

		x = p.x - playerSpeed
		if checkBoundaryHit(x, p.imageWidth) {
			return
		}
		p.x = x
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		x = p.x + playerSpeed
		if checkBoundaryHit(x, p.imageWidth) {
			return
		}
		p.x = x
	}
}
func checkBoundaryHit(x float64, w int32) bool {
	return x < 0 || x > float64(screenWidth-w)
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
	p.imageHeight = imageHeight
	p.imageWidth = imageWidth
	p.Tex = playerTexture
	p.x = screenWidth/2.0 - float64(imageWidth/2.0)
	p.y = screenHeight - float64(imageHeight)
	return p, nil
}

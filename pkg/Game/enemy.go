package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Enemy struct {
	Tex                     *sdl.Texture
	x, y                    float64
	imageHeight, imageWidth int32
}

func (e *Enemy) Draw() {
	Renderer.CopyEx(e.Tex, &sdl.Rect{X: 0, Y: 0, W: e.imageWidth, H: e.imageHeight}, &sdl.Rect{
		X: 0,
		Y: 0,
		W: e.imageWidth,
		H: e.imageHeight,
	}, 180.0,
		&sdl.Point{X: e.imageWidth / 2, Y: e.imageHeight / 2},
		sdl.FLIP_NONE)
}
func NewEnemy() (Enemy, error) {
	var e Enemy
	surface, err := sdl.LoadBMP("sprites/basic_enemy.bmp")
	if err != nil {
		return Enemy{}, fmt.Errorf("loading enemy %v", err)
	}
	defer surface.Free()
	imageWidth := surface.W
	imageHeight := surface.H
	e.imageWidth = imageWidth
	e.imageHeight = imageHeight
	enemyTex, err := Renderer.CreateTextureFromSurface(surface)
	e.Tex = enemyTex
	if err != nil {
		return Enemy{}, fmt.Errorf("creating texture", err)
	}
	if err != nil {
		return Enemy{}, fmt.Errorf("copying texture to renderer", err)
	}
	return e, nil

}

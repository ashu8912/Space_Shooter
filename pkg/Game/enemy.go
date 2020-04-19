package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySize = 105
)
const (
	EnemyXGridCount = 5
	EnemyYGridCount = 3
)

type Enemy struct {
	Tex                     *sdl.Texture
	x, y                    float64
	imageHeight, imageWidth int32
}

func (e *Enemy) Draw() {
	Renderer.CopyEx(e.Tex, &sdl.Rect{X: 0, Y: 0, W: e.imageWidth, H: e.imageHeight}, &sdl.Rect{
		X: (int32)(e.x),
		Y: (int32)(e.y),
		W: e.imageWidth,
		H: e.imageHeight,
	}, 180.0,
		&sdl.Point{X: e.imageWidth / 2, Y: e.imageHeight / 2},
		sdl.FLIP_NONE)
}
func NewEnemy(x, y float64) (Enemy, error) {
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
	e.x = x
	e.y = y
	if err != nil {
		return Enemy{}, fmt.Errorf("creating texture", err)
	}
	if err != nil {
		return Enemy{}, fmt.Errorf("copying texture to renderer", err)
	}
	return e, nil

}

func CreateMultipleEnemy(e []Enemy) []Enemy {
	for i := 0; i < EnemyXGridCount; i++ {
		for j := 0; j < EnemyYGridCount; j++ {
			x := float64(i) / 5 * screenWidth
			y := float64(j) * basicEnemySize
			enemy, _ := NewEnemy(x, y)
			e = append(e, enemy)
		}
	}
	return e
}

package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const bulletSpeed = 0.5

type Bullet struct {
	Tex         *sdl.Texture
	imageWidth  int32
	imageHeight int32
	x, y        float64
	active      bool
}

func (b *Bullet) Draw() {
	if !b.active {
		return
	}
	Renderer.Copy(b.Tex, &sdl.Rect{
		X: 0,
		Y: 0,
		W: b.imageWidth,
		H: b.imageHeight,
	}, &sdl.Rect{
		X: (int32)(b.x),
		Y: (int32)(b.y),
		W: b.imageWidth,
		H: b.imageHeight,
	})
}
func NewBullet() (*Bullet, error) {
	var b Bullet
	surface, err := sdl.LoadBMP("sprites/player_bullet.bmp")
	defer surface.Free()
	if err != nil {
		return nil, fmt.Errorf("loading bullet image %v", err)
	}
	imageWidth := surface.W
	imageHeight := surface.H
	bulletTexture, err := Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, fmt.Errorf("creatig bullet texture %v", err)
	}
	b.imageHeight = imageHeight
	b.imageWidth = imageWidth
	b.Tex = bulletTexture
	b.x = screenWidth/2.0 - float64(imageWidth)
	b.y = screenHeight - float64(imageHeight)
	return &b, nil
}
func (b *Bullet) Update() {
	if b.y < 0 {
		b.active = false
		return
	}
	b.y -= bulletSpeed
}

var BulletPool []*Bullet

func InitBulletPool() {
	for i := 0; i < 30; i++ {
		bul, _ := NewBullet()
		BulletPool = append(BulletPool, bul)
	}
}

func bulletFromPool() (*Bullet, bool) {
	for _, bul := range BulletPool {
		if bul != nil {
			if !bul.active {
				return bul, true
			}
		}

	}
	return nil, false
}

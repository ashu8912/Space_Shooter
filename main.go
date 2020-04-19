package main

import (
	"log"

	game "github.com/ashu8912/gaming/pkg/Game"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	game.InitWindow()
	defer game.Window.Destroy()
	game.InitRenderer()
	defer game.Renderer.Destroy()

	running := true
	var err error
	var p game.Player
	p, err = game.NewPlayer()
	var e []game.Enemy
	enemies := game.CreateMultipleEnemy(e)
	for i := 0; i < game.EnemyXGridCount*game.EnemyYGridCount; i++ {
		defer enemies[i].Tex.Destroy()
	}
	defer p.Tex.Destroy()
	if err != nil {
		log.Fatal(err)
		return
	}
	game.InitBulletPool()
	for running {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		game.Renderer.SetDrawColor(255, 255, 255, 255)
		game.Renderer.Clear()
		for i := 0; i < game.EnemyXGridCount*game.EnemyYGridCount; i++ {
			enemies[i].Draw()
		}
		p.Draw()
		for _, bul := range game.BulletPool {
			if bul != nil {
				bul.Draw()
				bul.Update()
			}

		}
		p.Update()
		game.Renderer.Present()
	}
}

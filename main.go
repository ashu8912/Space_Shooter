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
	player, err := game.NewPlayer()
	defer player.Tex.Destroy()
	if err != nil {
		log.Fatal(err)
		return
	}
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

		game.Renderer.Copy(player.Tex, &sdl.Rect{
			X: 0,
			Y: 0,
			W: player.ImageWidth,
			H: player.ImageHeight,
		}, &sdl.Rect{
			X: 0,
			Y: 0,
			W: player.ImageWidth,
			H: player.ImageHeight,
		})
		game.Renderer.Present()
	}
}

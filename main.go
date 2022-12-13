package main

import (
	twenty48 "github.com/RaRseR/2048/2048"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	game, err := twenty48.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(twenty48.ScreenWidth, twenty48.ScreenHeight)
	ebiten.SetWindowTitle("github.com/RaRseR/2048")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

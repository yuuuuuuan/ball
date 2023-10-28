package main

import (
	ballgame "ball/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	game, err := ballgame.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(ballgame.ScreenWidth, ballgame.ScreenHeight)
	ebiten.SetWindowTitle("2048")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

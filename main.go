package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/eozgit/life/game"
	"github.com/eozgit/life/theme"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(game.Width*8, game.Height*8)
	ebiten.SetWindowTitle("Life")
	game := game.Game{Speed: game.DefaultSpeed, Theme: &theme.BlackAndWhite{}}
	game.ResetTiles(.2)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

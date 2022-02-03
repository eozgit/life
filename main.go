package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var width, height = 160, 120

var defaultSpeed = 5

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(width*8, height*8)
	ebiten.SetWindowTitle("Life")
	shouldIterate := makeShouldIterate()
	game := Game{0, 0, nil, defaultSpeed, shouldIterate, false}
	game.resetTiles(.2)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var numberKeys = []ebiten.Key{ebiten.Key0, ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9}

var colourFuncMap = map[ebiten.Key]func(cell *Cell, age int) color.RGBA{
	ebiten.Key1: blackAndWhite,
	ebiten.Key2: cga,
	ebiten.Key3: earth,
}

func (g *Game) checkInput() {
	if ebiten.IsKeyPressed(ebiten.KeyT) {
		for k, v := range colourFuncMap {
			if inpututil.IsKeyJustPressed(k) {
				g.colour = v
			}
		}
		return
	}

	for _, key := range numberKeys {
		if inpututil.IsKeyJustPressed(key) {
			g.speed = int(key) - int(ebiten.Key0)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		if g.showHelp {
			g.speed = defaultSpeed
		} else {
			g.speed = 0
		}
		g.showHelp = !g.showHelp
	}
}

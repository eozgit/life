package game

import (
	"github.com/eozgit/life/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var numberKeys = []ebiten.Key{ebiten.Key0, ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9}

var colourFuncMap = map[ebiten.Key]theme.Theme{
	ebiten.Key1: &theme.BlackAndWhite{},
	ebiten.Key2: &theme.CGA{},
	ebiten.Key3: &theme.Earth{},
}

func (g *Game) checkInput() {
	if ebiten.IsKeyPressed(ebiten.KeyT) {
		for k, v := range colourFuncMap {
			if inpututil.IsKeyJustPressed(k) {
				g.Theme = v
			}
		}
		return
	}

	for _, key := range numberKeys {
		if inpututil.IsKeyJustPressed(key) {
			g.Speed = int(key) - int(ebiten.Key0)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		if g.ShowHelp {
			g.Speed = DefaultSpeed
		} else {
			g.Speed = 0
		}
		g.ShowHelp = !g.ShowHelp
	}
}

package game

import (
	"log"

	"github.com/eozgit/life/game/cell"
	"github.com/eozgit/life/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var numberKeys = []ebiten.Key{ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9}

func (g *Game) checkInput() {
	if ebiten.IsKeyPressed(ebiten.KeyT) {
		for k, v := range theme.ThemeMap {
			if inpututil.IsKeyJustPressed(k) {
				g.Theme = v
				log.Printf("Select theme %d", k-43)
			}
		}
		return
	}

	for _, key := range numberKeys {
		if inpututil.IsKeyJustPressed(key) {
			g.Speed = int(key) - int(ebiten.Key0)
			log.Printf("Set speed to %d", g.Speed)
			return
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.ShowHelp = !g.ShowHelp
		if g.ShowHelp {
			log.Printf("Show help")
		} else {
			log.Printf("Hide help")
		}
		return
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.Cells[x][y] = cell.Cell{Alive: true, Iteration: g.Iteration - 1}
		if _, ok := g.Changes[x]; !ok {
			g.Changes[x] = make(map[int]struct{})
		}
		g.Changes[x][y] = struct{}{}
		return
	}
}

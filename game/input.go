package game

import (
	"log"

	"github.com/eozgit/life/game/cell"
	"github.com/eozgit/life/pattern"
	"github.com/eozgit/life/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var numberKeys = []ebiten.Key{ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9}

var lastSpaceshipCreatedAt = 0

func (g *Game) handleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		for _, key := range numberKeys {
			if inpututil.IsKeyJustPressed(key) {
				density := float32(int(key)-int(ebiten.Key0)) / 10
				g.ResetTiles(density)
				log.Printf("Reset %.1f", density)
				return
			}
		}
	}

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
		if lastSpaceshipCreatedAt > g.Iteration-g.Speed {
			return
		}

		x, y := ebiten.CursorPosition()
		if ebiten.IsKeyPressed(ebiten.KeyZ) {
			g.resurrectByPattern("block", x, y)
			log.Printf("Create block at %d, %d", x, y)
		} else if ebiten.IsKeyPressed(ebiten.KeyX) {
			g.createSpaceship(x, y, "glider", "glider")
		} else if ebiten.IsKeyPressed(ebiten.KeyC) {
			g.createSpaceship(x, y, "lwss", "light-weight spaceship")
		} else if ebiten.IsKeyPressed(ebiten.KeyV) {
			g.createSpaceship(x, y, "mwss", "middle-weight spaceship")
		} else if ebiten.IsKeyPressed(ebiten.KeyB) {
			g.createSpaceship(x, y, "hwss", "heavy-weight spaceship")
		} else {
			g.resurrectByPattern("cell", x, y)
			log.Printf("Resurrect cell %d, %d", x, y)
		}
		return
	}
}

func (g *Game) createSpaceship(x int, y int, patternName string, spaceshipName string) {
	lastSpaceshipCreatedAt = g.Iteration
	g.resurrectByPattern(patternName, x, y)
	log.Printf("Create %s at %d, %d", spaceshipName, x, y)
}

func (g *Game) resurrectByPattern(patternName string, x int, y int) {
	coords := pattern.GetPattern(patternName)
	for i := range coords {
		for j := range coords[i] {
			g.resurrectCell(x+i, y+j)
		}
	}
}

func (g *Game) resurrectCell(x int, y int) {
	if x >= Width {
		x %= Width
	}
	if y >= Height {
		y %= Height
	}
	g.Cells[x][y] = cell.Cell{Alive: true, Iteration: g.Iteration - 1}
	if _, ok := g.Changes[x]; !ok {
		g.Changes[x] = make(map[int]struct{})
	}
	g.Changes[x][y] = struct{}{}
}

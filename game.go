package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Cell struct {
	alive     bool
	iteration int
}

type Game struct {
	tick                int
	iteration           int
	cells               map[int]map[int]Cell
	speed               int
	shouldIterateCached func(speed int, tick int) bool
	showHelp            bool
	changes             map[int]map[int]struct{}
}

func (g *Game) Update() error {
	g.checkInput()

	if g.shouldIterate() {
		g.iterate()
		g.iteration++
	}
	g.tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.showHelp {
		ebitenutil.DebugPrint(screen, "0-9 set speed\nh   resume")
		return
	}

	screen.Fill(color.White)
	g.scan(func(i int, j int) {
		if g.cells[i][j].alive {
			screen.Set(i, j, color.Black)
		}
	})

	pink, aqua := color.RGBA{255, 0, 255, 255}, color.RGBA{0, 255, 255, 255}

	g.scanChanges(func(i int, j int) {
		var colour color.RGBA
		if g.cells[i][j].alive {
			colour = pink
		} else {
			colour = aqua
		}
		screen.Set(i, j, colour)
	})

	if g.tick < 80 {
		g.scan(func(i int, j int) {
			if j < 19 {
				screen.Set(i, j, color.Black)
			}
		})

		ebitenutil.DebugPrint(screen, "hit h for help")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func (g *Game) resetTiles(density float32) {
	g.cells = make(map[int]map[int]Cell)
	g.scan(func(i int, j int) {
		if g.cells[i] == nil {
			g.cells[i] = make(map[int]Cell)
		}
		alive := rand.Float32() < density
		g.cells[i][j] = Cell{alive, 0}
	})
	g.changes = make(map[int]map[int]struct{})
	g.scan(func(i int, j int) {
		if g.changes[i] == nil {
			g.changes[i] = make(map[int]struct{})
		}
		g.changes[i][j] = struct{}{}
	})
}

func (g *Game) shouldIterate() bool {
	return g.shouldIterateCached(g.speed, g.tick)
}

var numberKeys = []ebiten.Key{ebiten.Key0, ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9}

func (g *Game) checkInput() {
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

func (g *Game) iterate() {
	potentialChanges := g.getPotentialChanges()

	changes := g.getChanges(potentialChanges)

	for i := range changes {
		for j := range changes[i] {
			g.cells[i][j] = Cell{!g.cells[i][j].alive, g.iteration}
		}
	}

	g.changes = changes
}

func (g *Game) getAliveCountWithinProximity(i int, j int) int {
	alive := 0
neighbourhood:
	for m := 0; m < 3; m++ {
		im := i + m - 1
		if !(im < 0) && im < width {
			for n := 0; n < 3; n++ {
				jn := j + n - 1
				if !(jn < 0) && jn < height {
					if g.cells[im][jn].alive {
						alive++
					}
					if alive > 4 {
						break neighbourhood
					}
				}
			}
		}
	}
	return alive
}

func (g *Game) getPotentialChanges() map[int]map[int]struct{} {
	potentialChanges := make(map[int]map[int]struct{})
	g.scanChanges(func(i int, j int) {
		for m := 0; m < 3; m++ {
			im := i + m - 1
			if !(im < 0) && im < width {
				for n := 0; n < 3; n++ {
					jn := j + n - 1
					if !(jn < 0) && jn < height {
						if _, ok := potentialChanges[im]; !ok {
							potentialChanges[im] = make(map[int]struct{})
						}
						potentialChanges[im][jn] = struct{}{}
					}
				}
			}
		}
	})
	return potentialChanges
}

func (g *Game) getChanges(potentialChanges map[int]map[int]struct{}) map[int]map[int]struct{} {
	changes := make(map[int]map[int]struct{})
	for i := range potentialChanges {
		for j := range potentialChanges[i] {
			alive := g.getAliveCountWithinProximity(i, j)

			if alive == 3 && !g.cells[i][j].alive {
				if _, ok := changes[i]; !ok {
					changes[i] = make(map[int]struct{})
				}
				changes[i][j] = struct{}{}
			}

			if (alive < 3 || alive > 4) && g.cells[i][j].alive {
				if _, ok := changes[i]; !ok {
					changes[i] = make(map[int]struct{})
				}
				changes[i][j] = struct{}{}
			}
		}
	}
	return changes
}

func (g *Game) scan(callback func(i int, j int)) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			callback(i, j)
		}
	}
}

func (g *Game) scanChanges(callback func(i int, j int)) {
	for i := range g.changes {
		for j := range g.changes[i] {
			callback(i, j)
		}
	}
}

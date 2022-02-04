package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Tile struct {
	alive     bool
	iteration int
}

type Game struct {
	tick                int
	iteration           int
	tiles               map[int]map[int]Tile
	speed               int
	shouldIterateCached func(speed int, tick int) bool
	help                bool
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
	if g.help {
		ebitenutil.DebugPrint(screen, "0-9 set speed\nh   resume")
	} else {
		screen.Fill(color.White)
		g.sweep(func(i int, j int) {
			if g.tiles[i][j].alive {
				screen.Set(i, j, color.Black)
			}
		})

		if g.tick < 150 {
			g.sweep(func(i int, j int) {
				if j < 19 {
					screen.Set(i, j, color.Black)
				}
			})

			ebitenutil.DebugPrint(screen, "hit h for help")
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func (g *Game) resetTiles(density float32) {
	g.tiles = make(map[int]map[int]Tile)
	g.sweep(func(i int, j int) {
		if g.tiles[i] == nil {
			g.tiles[i] = make(map[int]Tile)
		}
		alive := rand.Float32() < density
		g.tiles[i][j] = Tile{alive, 0}
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
		if g.help {
			g.speed = defaultSpeed
		} else {
			g.speed = 0
		}
		g.help = !g.help
	}
}

func (g *Game) iterate() {
	updatedLastRound := make(map[int]map[int]struct{})

	g.sweep(func(i int, j int) {
		tile := g.tiles[i][j]
		if tile.iteration == g.iteration-1 {
			for m := 0; m < 3; m++ {
				im := i + m - 1
				if !(im < 0) && im < len(g.tiles) {
					for n := 0; n < 3; n++ {
						jn := j + n - 1
						if !(jn < 0) && jn < len(g.tiles[i]) {
							if _, ok := updatedLastRound[im]; !ok {
								updatedLastRound[im] = make(map[int]struct{})
							}
							updatedLastRound[im][jn] = struct{}{}
						}
					}
				}
			}
		}
	})

	updates := make(map[int]map[int]bool)
	for i := range updatedLastRound {
		for j := range updatedLastRound[i] {
			alive := 0
		neighbourhood:
			for m := 0; m < 3; m++ {
				im := i + m - 1
				if !(im < 0) && im < len(g.tiles) {
					for n := 0; n < 3; n++ {
						jn := j + n - 1
						if !(jn < 0) && jn < len(g.tiles[i]) {
							if g.tiles[im][jn].alive {
								alive++
							}
							if alive > 4 {
								break neighbourhood
							}
						}
					}
				}
			}

			if alive == 3 && !g.tiles[i][j].alive {
				if _, ok := updates[i]; !ok {
					updates[i] = make(map[int]bool)
				}
				updates[i][j] = true
			}

			if (alive < 3 || alive > 4) && g.tiles[i][j].alive {
				if _, ok := updates[i]; !ok {
					updates[i] = make(map[int]bool)
				}
				updates[i][j] = false
			}
		}
	}

	for i := range updates {
		for j := range updates[i] {
			g.tiles[i][j] = Tile{updates[i][j], g.iteration}
		}
	}
}

func (g *Game) sweep(callback func(i int, j int)) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			callback(i, j)
		}
	}
}

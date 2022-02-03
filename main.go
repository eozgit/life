package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var dim = 100

type Tile struct {
	alive     bool
	iteration int
}

type Game struct {
	tick                int
	iteration           int
	tiles               [][]Tile
	speed               int
	shouldIterateCached func(speed int, tick int) bool
}

func (g *Game) Update() error {
	g.checkInput()

	if g.shouldIterate() {
		g.resetTiles(0.2)
	}
	g.tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for i := range g.tiles {
		for j := range g.tiles[i] {
			if g.tiles[i][j].alive {
				screen.Set(i, j, color.Black)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return dim, dim
}

func (g *Game) resetTiles(density float32) {
	g.tiles = make([][]Tile, dim)
	for i := range g.tiles {
		g.tiles[i] = make([]Tile, dim)
		for j := 0; j < dim; j++ {
			alive := rand.Float32() < density
			g.tiles[i][j] = Tile{alive, 0}
		}
	}
}

func (g *Game) shouldIterate() bool {
	return g.shouldIterateCached(g.speed, g.tick%1000)
}

var numberKeys = []ebiten.Key{ebiten.Key0, ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9}

func (g *Game) checkInput() {
	for _, key := range numberKeys {
		if inpututil.IsKeyJustPressed(key) {
			g.speed = int(key) - int(ebiten.Key0)
		}
	}
}

var decay = .5

func makeShouldIterate() func(speed int, tick int) bool {
	speedTickIterationMap := make(map[int]map[int]int)
	for i := 1; i < 10; i++ {
		speedTickIterationMap[i] = make(map[int]int)
	}
	return func(speed int, tick int) bool {
		if speed == 0 {
			return false
		}

		tickIterationMap := speedTickIterationMap[speed]
		var iteration int
		if iter, ok := tickIterationMap[tick]; ok {
			iteration = iter
		} else {
			factor := math.Pow(decay, float64(9-speed))
			iteration = int(float64(tick) * factor)
			tickIterationMap[tick] = iteration
		}

		prev := tickIterationMap[tick-1]
		return iteration > prev
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(dim*8, dim*8)
	ebiten.SetWindowTitle("Life")
	shouldIterate := makeShouldIterate()
	game := Game{0, 0, nil, 1, shouldIterate}
	game.resetTiles(0.2)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

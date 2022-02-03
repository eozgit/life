package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var width, height = 160, 120

var defaultSpeed = 5

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
	help                bool
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
	if g.help {
		ebitenutil.DebugPrint(screen, "0-9 set speed\nh   resume")
	} else {
		screen.Fill(color.White)
		for i := range g.tiles {
			for j := range g.tiles[i] {
				if g.tiles[i][j].alive {
					screen.Set(i, j, color.Black)
				}
			}
		}
		if g.tick < 150 {
			for i := 0; i < width; i++ {
				for j := 0; j < 19; j++ {
					screen.Set(i, j, color.Black)
				}
			}
			ebitenutil.DebugPrint(screen, "hit h for help")
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func (g *Game) resetTiles(density float32) {
	g.tiles = make([][]Tile, width)
	for i := range g.tiles {
		g.tiles[i] = make([]Tile, height)
		for j := 0; j < height; j++ {
			alive := rand.Float32() < density
			g.tiles[i][j] = Tile{alive, 0}
		}
	}
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

var decay = .5

func makeShouldIterate() func(speed int, tick int) bool {
	speedTickIterationMap := make(map[int]map[int]int)
	for i := 1; i < 10; i++ {
		speedTickIterationMap[i] = make(map[int]int)
	}
	mode := 1000
	return func(speed int, tick int) bool {
		step := tick % mode

		if speed == 0 {
			return false
		}

		tickIterationMap := speedTickIterationMap[speed]
		var iteration int
		if iter, ok := tickIterationMap[step]; ok {
			iteration = iter
		} else {
			factor := math.Pow(decay, float64(9-speed))
			iteration = int(float64(step) * factor)
			tickIterationMap[step] = iteration
		}

		idx := mode - 1
		if step > 0 {
			idx = step - 1
		}
		prev := tickIterationMap[idx]
		return iteration > prev
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(width*8, height*8)
	ebiten.SetWindowTitle("Life")
	shouldIterate := makeShouldIterate()
	game := Game{0, 0, nil, defaultSpeed, shouldIterate, false}
	game.resetTiles(0.2)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

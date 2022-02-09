package game

import (
	"image/color"
	"math/rand"

	"github.com/eozgit/life/game/cell"
	"github.com/eozgit/life/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Tick      int
	Iteration int
	Speed     int
	Cells     map[int]map[int]cell.Cell
	Changes   map[int]map[int]struct{}
	ShowHelp  bool
	Theme     theme.Theme
}

func (g *Game) Update() error {
	g.checkInput()

	if g.shouldIterate() {
		g.iterate()
		g.Iteration++
	}
	g.Tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.ShowHelp {
		ebitenutil.DebugPrint(screen, "0-9      set speed\nt + 1-4  set theme\nh        resume")
		return
	}

	screen.Fill(color.White)
	g.scan(func(i int, j int) {
		cell := g.Cells[i][j]
		age := g.Iteration - cell.Iteration - 1
		p := &cell
		colour := g.Theme.Colour(p, age)
		screen.Set(i, j, colour)
	})

	if g.Tick < 80 {
		g.scan(func(i int, j int) {
			if j < 19 {
				screen.Set(i, j, color.Black)
			}
		})

		ebitenutil.DebugPrint(screen, "hit h for help")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func (g *Game) ResetTiles(density float32) {
	g.Cells = make(map[int]map[int]cell.Cell)
	g.scan(func(i int, j int) {
		if g.Cells[i] == nil {
			g.Cells[i] = make(map[int]cell.Cell)
		}
		alive := rand.Float32() < density
		g.Cells[i][j] = cell.Cell{Alive: alive, Iteration: 0}
	})
	g.Changes = make(map[int]map[int]struct{})
	g.scan(func(i int, j int) {
		if g.Changes[i] == nil {
			g.Changes[i] = make(map[int]struct{})
		}
		g.Changes[i][j] = struct{}{}
	})
}

func (g *Game) shouldIterate() bool {
	return !g.ShowHelp && ShouldIterate(g.Speed, g.Tick)
}

func (g *Game) iterate() {
	potentialChanges := g.getPotentialChanges()

	changes := g.getChanges(potentialChanges)

	for i := range changes {
		for j := range changes[i] {
			g.Cells[i][j] = cell.Cell{Alive: !g.Cells[i][j].Alive, Iteration: g.Iteration}
		}
	}

	g.Changes = changes
}

func (g *Game) getPotentialChanges() map[int]map[int]struct{} {
	potentialChanges := make(map[int]map[int]struct{})
	g.scanChanges(func(i int, j int) {
		for m := 0; m < 3; m++ {
			im := i + m - 1
			if im < 0 {
				im += Width
			}
			im = im % Width

			for n := 0; n < 3; n++ {
				jn := j + n - 1
				if jn < 0 {
					jn += Height
				}
				jn = jn % Height

				if _, ok := potentialChanges[im]; !ok {
					potentialChanges[im] = make(map[int]struct{})
				}
				potentialChanges[im][jn] = struct{}{}
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

			if alive == 3 && !g.Cells[i][j].Alive {
				if _, ok := changes[i]; !ok {
					changes[i] = make(map[int]struct{})
				}
				changes[i][j] = struct{}{}
			}

			if (alive < 3 || alive > 4) && g.Cells[i][j].Alive {
				if _, ok := changes[i]; !ok {
					changes[i] = make(map[int]struct{})
				}
				changes[i][j] = struct{}{}
			}
		}
	}
	return changes
}

func (g *Game) getAliveCountWithinProximity(i int, j int) int {
	alive := 0
neighbourhood:
	for m := 0; m < 3; m++ {
		im := i + m - 1
		if im < 0 {
			im += Width
		}
		im = im % Width

		for n := 0; n < 3; n++ {
			jn := j + n - 1
			if jn < 0 {
				jn += Height
			}
			jn = jn % Height

			if g.Cells[im][jn].Alive {
				alive++
			}
			if alive > 4 {
				break neighbourhood
			}
		}
	}
	return alive
}

func (g *Game) scan(callback func(i int, j int)) {
	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {
			callback(i, j)
		}
	}
}

func (g *Game) scanChanges(callback func(i int, j int)) {
	for i := range g.Changes {
		for j := range g.Changes[i] {
			callback(i, j)
		}
	}
}

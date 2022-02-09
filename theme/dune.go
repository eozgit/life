package theme

import (
	"image/color"

	"github.com/eozgit/life/game/cell"
	"golang.org/x/image/colornames"
)

type Dune struct {
}

func (t *Dune) Colour(cell *cell.Cell, age int) color.RGBA {
	if cell.Alive {
		if age == 0 {
			return colornames.Goldenrod
		}
		if age < 3 {
			return colornames.Darkgoldenrod
		}

		mode := age / 3 % 2

		if mode == 0 {
			return colornames.Grey
		}
		return colornames.Dimgrey
	}

	if age == 0 {
		return colornames.Saddlebrown
	}

	if age < 17 {
		if age%5 < 3 {
			return colornames.Burlywood
		}
		return colornames.Tan
	}

	if age%7 < 4 {
		return colornames.Navajowhite
	}
	return colornames.Bisque
}

package theme

import (
	"image/color"

	"github.com/eozgit/life/game/cell"
	"golang.org/x/image/colornames"
)

type CGA struct {
}

func (t *CGA) Colour(cell *cell.Cell, age int) color.RGBA {
	change := age == 0
	if cell.Alive {
		if change {
			return colornames.Fuchsia
		}
		return colornames.Black
	}

	if change {
		return colornames.Aqua
	}

	return colornames.White
}

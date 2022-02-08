package theme

import (
	"image/color"

	"github.com/eozgit/life/game/cell"
	"golang.org/x/image/colornames"
)

type BlackAndWhite struct {
}

func (t *BlackAndWhite) Colour(cell *cell.Cell, age int) color.RGBA {
	if cell.Alive {
		return colornames.Black
	}

	return colornames.White
}

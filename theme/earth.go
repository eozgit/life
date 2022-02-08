package theme

import (
	"image/color"

	"github.com/eozgit/life/game/cell"
	"golang.org/x/image/colornames"
)

type Earth struct {
}

func (t *Earth) Colour(cell *cell.Cell, age int) color.RGBA {
	if cell.Alive {
		if age == 0 {
			return colornames.Yellow
		}

		mode := age / 3 % 2

		if mode == 0 {
			return colornames.Green
		}
		return colornames.Darkgreen
	}

	if age == 0 {
		return colornames.Deepskyblue
	}

	if age < 17 {
		if age%5 < 3 {
			return colornames.Blue
		}
		return colornames.Mediumblue
	}

	if age%7 < 4 {
		return colornames.Darkblue
	}
	return colornames.Navy
}

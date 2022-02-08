package theme

import (
	"image/color"

	"github.com/eozgit/life/game/cell"
)

type Theme interface {
	Colour(cell *cell.Cell, age int) color.RGBA
}

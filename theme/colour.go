package theme

import (
	"image/color"

	"github.com/eozgit/life/game/cell"
	"github.com/hajimehoshi/ebiten/v2"
)

type Theme interface {
	Colour(cell *cell.Cell, age int) color.RGBA
}

var ThemeMap = map[ebiten.Key]Theme{
	ebiten.Key1: &BlackAndWhite{},
	ebiten.Key2: &CGA{},
	ebiten.Key3: &Earth{},
	ebiten.Key4: &Dune{},
}

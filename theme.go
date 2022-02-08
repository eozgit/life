package main

import (
	"image/color"

	"golang.org/x/image/colornames"
)

func blackAndWhite(cell *Cell, age int) color.RGBA {
	if cell.alive {
		return colornames.Black
	}

	return colornames.White
}

func cga(cell *Cell, age int) color.RGBA {
	change := age == 0
	if cell.alive {
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

func earth(cell *Cell, age int) color.RGBA {
	if cell.alive {
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

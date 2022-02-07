package main

import "image/color"

var black, white, pink, aqua = color.RGBA{0, 0, 0, 255}, color.RGBA{255, 255, 255, 255}, color.RGBA{255, 0, 255, 255}, color.RGBA{0, 255, 255, 255}

func blackAndWhite(cell *Cell, change bool) color.RGBA {
	if cell.alive {
		return black
	}

	return white
}

func cga(cell *Cell, change bool) color.RGBA {
	if cell.alive {
		if change {
			return pink
		}
		return black
	}

	if change {
		return aqua
	}

	return white
}

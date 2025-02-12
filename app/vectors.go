package app

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawVectors(screen *ebiten.Image) {
	vector.StrokeCircle(screen, 960, 900, 350, 10, color.Black, true)
	vector.StrokeCircle(screen, 210, 250, 150, 10, color.Black, true)
	vector.StrokeCircle(screen, 1670, 250, 200, 10, color.Black, true)
}

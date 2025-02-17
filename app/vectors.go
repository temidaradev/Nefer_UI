package app

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawVectors(screen *ebiten.Image) {
	vector.StrokeCircle(screen, float32(width)/2, float32(height)-100, 350, 10, color.Black, true)
	vector.StrokeCircle(screen, float32(width)/2-550, float32(height)-740, 150, 10, color.Black, true)
	vector.StrokeCircle(screen, float32(width)/2+500, float32(height)-740, 200, 10, color.Black, true)
}

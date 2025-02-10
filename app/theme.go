package app

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type DarkTheme struct {
	Background color.RGBA
}

func (dt *DarkTheme) SetTheme(screen *ebiten.Image, color color.RGBA) {
	screen.Fill(color)
}

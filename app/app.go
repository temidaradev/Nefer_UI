package app

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type UI struct {
	theme *DarkTheme
}

var (
	bkgColor = color.RGBA{95, 95, 95, 0xff}
	width    = 1280
	height   = 720
)

func MakeUI() *UI {
	u := &UI{}

	return u
}

func (u *UI) Draw(screen *ebiten.Image) {
	u.theme.SetTheme(screen, bkgColor)
}

func (u *UI) Update() error {
	return nil
}

func (u *UI) Layout(screenWidth, screenHeight int) (int, int) {
	return width, height
}

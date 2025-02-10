package app

import (
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/temidaradev/esset/v2"
)

func init() {
	assets.FontFace, _ = esset.GetFont(assets.TangerineRegular, 48)
}

type App struct {
	ui    *Ui
	theme *DarkTheme
}

var (
	bkgColor      = color.RGBA{95, 95, 95, 0xff}
	width, height = ebiten.Monitor().Size()
)

func MakeApp() *App {
	u := &App{
		ui: &Ui{},
	}

	return u
}

func (a *App) Draw(screen *ebiten.Image) {
	a.theme.SetTheme(screen, bkgColor)
	vector.StrokeCircle(screen, 960, 900, 350, 10, color.Black, true)
	esset.DrawText(screen, "Cahit", 48, float64(width)-150, float64(height)-150, assets.FontFace, color.White)
}

func (a *App) Update() error {
	return nil
}

func (a *App) Layout(screenWidth, screenHeight int) (int, int) {
	return screenWidth, screenHeight
}

package app

import (
	"fmt"
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/esset/v2"
)

func WriteData(screen *ebiten.Image) {
	vel := fmt.Sprintf("%v", 20)
	esset.DrawText(screen, "Km/s", 24, float64(width)/2-45, float64(height)/2+55, assets.FontFaceH, color.White)
	esset.DrawText(screen, vel, 64, float64(width)/2-50, float64(height)/2+250, assets.FontFaceT, color.White)
	thrm := fmt.Sprintf("%v", 30)
	esset.DrawText(screen, "C°", 24, float64(width)/2-565, float64(height)/2-400, assets.FontFaceH, color.White)
	esset.DrawText(screen, thrm, 48, float64(width)/2-585, float64(height)/2-300, assets.FontFaceV, color.White)
}

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
	esset.DrawText(screen, "Km/s", 24, float64(width)/2-130, float64(height)/2+250, assets.FontFaceV, color.White)
	esset.DrawText(screen, vel, 48, float64(width)/2-130, float64(height)/2+250, assets.FontFaceH, color.White)
	thrm := fmt.Sprintf("C° %v", 30)
	esset.DrawText(screen, thrm, 24, float64(width)/2-230, float64(height)/2-250, assets.FontFaceV, color.White)
}

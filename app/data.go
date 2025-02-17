package app

import (
	"fmt"
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/esset/v2"
)

func DrawLine(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(width)/2+370, float64(height)/2-400)

	screen.DrawImage(assets.Line, op)
}

func WriteData(screen *ebiten.Image) {
	vel := fmt.Sprintf("%v", 20)
	esset.DrawText(screen, "Km/s", 24, float64(width)/2-45, float64(height)/2+55, assets.FontFaceH, color.White)
	esset.DrawText(screen, vel, 64, float64(width)/2-50, float64(height)/2+250, assets.FontFaceT, color.White)

	thrm := fmt.Sprintf("%v", 30)
	esset.DrawText(screen, "C°", 24, float64(width)/2-565, float64(height)/2-400, assets.FontFaceH, color.White)
	esset.DrawText(screen, thrm, 48, float64(width)/2-585, float64(height)/2-300, assets.FontFaceV, color.White)

	volt := fmt.Sprintf("%v", 40)
	esset.DrawText(screen, " V", 24, float64(width)/2+485, float64(height)/2-450, assets.FontFaceH, color.White)
	esset.DrawText(screen, volt, 48, float64(width)/2+465, float64(height)/2-375, assets.FontFaceV, color.White)

	DrawLine(screen)
	watt := fmt.Sprintf("%v", 50)
	esset.DrawText(screen, "Wh", 24, float64(width)/2+480, float64(height)/2-260, assets.FontFaceH, color.White)
	esset.DrawText(screen, watt, 48, float64(width)/2+470, float64(height)/2-190, assets.FontFaceV, color.White)
}

package app

import (
	"fmt"
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/esset/v2"
)

func WriteData(screen *ebiten.Image) {
	vel := fmt.Sprintf("Km/s: %v")
	esset.DrawText(screen, vel, 24, 300, 500, assets.FontFaceV, color.White)

}

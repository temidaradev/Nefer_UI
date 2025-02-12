package app

import (
	"main/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Needle struct {
	angle float64
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

func (n *Needle) DrawNeedle(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(64)/2, -float64(64)/2)
	//op.GeoM.Rotate(DegreesToRadians(n.angle))
	op.GeoM.Scale(0.7, 0.7)
	op.GeoM.Translate(200, 400)

	screen.DrawImage(assets.Line, op)
}

func (n *Needle) UpdateNeedle() {

}

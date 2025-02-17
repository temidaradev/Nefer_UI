package assets

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/temidaradev/esset/v2"
)

//go:embed *
var Assets embed.FS

//go:embed fonts/Tangerine-Regular.ttf
var TangerineRegular []byte
var FontFace text.Face

//go:embed fonts/Poppins-Black.ttf
var PoppinsBlack []byte
var FontFaceV text.Face
var FontFaceH text.Face

var Line = esset.GetAsset(Assets, "images/needle.png")

package assets

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed *
var Assets embed.FS

//go:embed fonts/Tangerine-Regular.ttf
var TangerineRegular []byte
var FontFace text.Face

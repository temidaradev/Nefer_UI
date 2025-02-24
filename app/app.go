package app

import (
	"fmt"
	"image/color"
	"main/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/esset/v2"
)

func init() {
	assets.FontFace, _ = esset.GetFont(assets.TangerineRegular, 48)
	assets.FontFaceV, _ = esset.GetFont(assets.PoppinsBlack, 48)
	assets.FontFaceH, _ = esset.GetFont(assets.PoppinsBlack, 24)
	assets.FontFaceT, _ = esset.GetFont(assets.PoppinsBlack, 64)
	assets.FontFaceS, _ = esset.GetFont(assets.PoppinsBlack, 16)
}

type App struct {
	theme *DarkTheme
}

var (
	bkgColor      = color.RGBA{95, 95, 95, 0xff}
	width, height = ebiten.Monitor().Size()
)

func MakeApp() *App {
	u := &App{}

	return u
}

func (a *App) Draw(screen *ebiten.Image) {
	a.theme.SetTheme(screen, bkgColor)
	DrawVectors(screen)
	th := time.Now()
	tf := th.Format("2006-01-02\n15:04:05.000")

	db := OpenDatabase()
	defer db.Close()

	values := FindAll(db)
	data := ReadCsvFile("./data/data.csv")
	InsertData(db, ParseData(data))

	for _, song := range values {
		fmt.Println(song)
	}

	WriteData(screen)
	esset.DrawText(screen, "Team Nefer", 48, float64(width)/2-90, float64(height)/2-450, assets.FontFace, color.White)
	esset.DrawText(screen, "Cahit", 48, float64(width)-150, float64(height)-150, assets.FontFace, color.White)
	t := fmt.Sprintf("%v", tf)
	esset.DrawText(screen, t, 16, float64(width)/2-650, float64(height)/2+300, assets.FontFaceS, color.White)
}

func (a *App) Update() error {
	return nil
}

func (a *App) Layout(screenWidth, screenHeight int) (int, int) {
	return screenWidth, screenHeight
}

package app

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"image/color"
	"main/assets"
	"strings"
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
func genRandForSession(byteLength int) (string, error) {
	// Create a byte slice to store the random bytes
	randomBytes := make([]byte, byteLength)

	// Read random bytes from crypto/rand
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Encode the random bytes to base64
	// Use URL-safe encoding to avoid characters that might cause issues in filenames
	encodedString := base64.URLEncoding.EncodeToString(randomBytes)

	// Remove any padding characters (=)
	encodedString = strings.ReplaceAll(encodedString, "=", "")

	return encodedString, nil
}

func (a *App) Draw(screen *ebiten.Image) {
	a.theme.SetTheme(screen, bkgColor)
	DrawVectors(screen)
	th := time.Now()
	tf := th.Format("2006-01-02\n15:04:05.000")

	// db := OpenDatabase()
	// defer db.Close()

	// values := FindAll(db)
	// data := ReadCsvFile("./data/data.csv")
	// InsertData(db, ParseData(data))

	// for _, song := range values {
	// 	fmt.Println(song)
	// }

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

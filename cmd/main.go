package main

import (
	"log"
	"main/app"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	a := app.MakeApp()

	ebiten.SetWindowTitle("Nefer UI")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(ebiten.Monitor().Size())

	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}
}

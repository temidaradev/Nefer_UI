package main

import (
	"log"
	"main/app"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	u := app.MakeUI()
	ebiten.SetWindowTitle("Nefer UI")
	if err := ebiten.RunGame(u); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"main/app"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	a := app.MakeApp()

	ebiten.SetWindowTitle("Nefer UI")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(ebiten.Monitor().Size())

	result := app.ReadCsvFile("./data/data.csv")
	fmt.Println(result)

	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}
}

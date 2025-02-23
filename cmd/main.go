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

	db := app.OpenDatabase()
	defer db.Close()

	data := app.ReadCsvFile("./data/data.csv")
	app.InsertData(db, app.ParseData(data))

	values := app.FindAll(db)

	for _, song := range values {
		fmt.Println(song)
	}

	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}
}

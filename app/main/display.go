package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/lennard-pfennig-bruker/screending/base"
	"image"
	"os"

	"github.com/fstanis/screenresolution"
)

func display(img image.Image) {
	a := app.New()
	w := a.NewWindow("Images")
	raw, _ := base.GetImageFromFilePath("/Users/karluwe/Pictures/cool/2010_Wolfing_Ostern.JPG")
	img := canvas.NewImageFromImage(raw)
	w.SetContent(img)

	resolution := screenresolution.GetPrimary()
	if resolution.String() == "" {
		fmt.Println("failed to get screen resolution")
		os.Exit(1)
	}
	fmt.Println(resolution)
	resolution.Width = 300
	resolution.Height = 400
	w.Resize(fyne.NewSize(float32(resolution.Width), float32(resolution.Height)))

	w.ShowAndRun()
}

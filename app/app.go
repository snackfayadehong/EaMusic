package myApp

import (
	"fyne.io/fyne/v2"
)

type AppConfig struct {
	App        fyne.App
	Height     float32
	Width      float32
	MainWindow fyne.Window
	TitleBar   fyne.Container
}

var MApp = AppConfig{Height: 670, Width: 1050}

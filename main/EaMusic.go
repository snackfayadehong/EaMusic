package main

import (
	myApp "EaMusic/app"
	"EaMusic/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	// application
	a := app.NewWithID("EaMusic")
	myApp.MApp.App = a
	myApp.MApp.App.Settings().SetTheme(&theme.MyTheme{})

	// window
	myApp.MApp.MainWindow = a.NewWindow("EaMusic")
	myApp.MApp.MainWindow.CenterOnScreen() // 居中
	myApp.MApp.MainWindow.Resize(fyne.Size{
		Width:  1000,
		Height: 600,
	})
	myApp.MApp.MainWindow.SetMaster() // 主窗口
	myApp.MApp.MainWindow.ShowAndRun()
}

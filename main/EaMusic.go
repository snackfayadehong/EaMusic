package main

import (
	myApp "EaMusic/app"
	"EaMusic/theme"
	"EaMusic/ui"
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
	// 无边框窗口
	// drv := fyne.CurrentApp().Driver()
	// if drv, ok := drv.(desktop.Driver); ok {
	// 	myApp.MApp.MainWindow = drv.CreateSplashWindow()
	// 	myApp.MApp.MainWindow.SetFullScreen(true)
	// }
	myApp.MApp.MainWindow.CenterOnScreen() // 居中
	myApp.MApp.MainWindow.Resize(fyne.Size{
		Width:  myApp.MApp.Width,
		Height: myApp.MApp.Height,
	})
	ui.MyUi.MainWindow = myApp.MApp.MainWindow
	myApp.MApp.MainWindow.SetMaster() // 主窗口
	ui.MyUi.MakeUi()
	myApp.MApp.MainWindow.ShowAndRun()

}

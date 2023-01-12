package main

import (
	myApp "EaMusic/app"
	"EaMusic/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// application
	// a := app.NewWithID("EaMusic")
	// myApp.MApp.App = a
	myApp.MApp.App.Settings().SetTheme(&theme.MyTheme{})

	// window
	// myApp.MApp.MainWindow = a.NewWindow("EaMusic")
	// 无边框窗口
	drv := fyne.CurrentApp().Driver()
	if drv, ok := drv.(desktop.Driver); ok {
		myApp.MApp.MainWindow = drv.CreateSplashWindow()
	}
	myApp.MApp.MainWindow.CenterOnScreen() // 居中
	myApp.MApp.MainWindow.Resize(fyne.Size{
		Width:  1050,
		Height: 670,
	})
	// myApp.MApp.MainWindow.SetMaster() // 主窗口
	// 添加标题栏
	titleBar := container.NewHBox(
		widget.NewLabel("EaMusic"),
	)
	myApp.MApp.MainWindow.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, titleBar, nil), titleBar))
	myApp.MApp.MainWindow.ShowAndRun()
}

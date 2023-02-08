package ui

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Ui struct {
	MainWindow fyne.Window
	TopBar     *fyne.Container // 顶部搜索栏
	MusicBar   *fyne.Container // 音乐信息栏
	PlayBar    *fyne.Container // 底部播放栏
}

var MyUi Ui

// SearchMusicBar 搜索框
func (u *Ui) SearchMusicBar() *fyne.Container {
	input := widget.NewEntry()
	input.PlaceHolder = "日落大道"
	inputContainer := container.NewGridWrap(fyne.NewSize(500, 35))
	inputContainer.Add(input)
	sh := container.NewHBox(inputContainer)
	return sh
}

// MusicDetailBar 音乐详细栏
// func (u *Ui) MusicDetailBar() *fyne.Container {
// 	var data [][]interface{}
// 	musicList := widget.NewList(
// 		func() int {
// 			return len(data)
// 		},
// 		func() fyne.CanvasObject {
// 			return container.NewVBox()
// 		},
// 		func(id widget.ListItemID, object fyne.CanvasObject) {
// 			object.(*fyne.Container).Objects = []fyne.CanvasObject{
// 				widget.NewLabel(),
// 			}
// 		},
// 	)
// }

// MakeUi 构建Ui
func (u *Ui) MakeUi() {
	search := container.NewHBox(layout.NewSpacer(), u.SearchMusicBar(), layout.NewSpacer())
	// search.Resize(fyne.NewSize(300, 20))
	// mainContent := container.NewVBox(layout.NewSpacer(),search, layout.NewSpacer())
	mainContent := container.NewVBox(search)
	u.MainWindow.SetContent(mainContent)
}

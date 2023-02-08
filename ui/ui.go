package ui

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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
func (u *Ui) MakeUi() {
	search := u.SearchMusicBar()
	// search.Resize(fyne.NewSize(300, 20))
	// mainContent := container.NewVBox(layout.NewSpacer(),search, layout.NewSpacer())
	mainContent := container.NewVBox(search)
	u.MainWindow.SetContent(mainContent)
}

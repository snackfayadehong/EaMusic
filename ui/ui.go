package ui

import (
	_ "embed"
	"fyne.io/fyne/v2"
)

type Ui struct {
	MainWindow fyne.Window
	FuncBar    *fyne.Container // 左侧功能栏
	TopBar     *fyne.Container // 顶部搜索栏
	MusicBar   *fyne.Container // 音乐信息栏
	PlayBar    *fyne.Container // 底部播放栏
}

var MyUi Ui

//go:embed EaMusic.png
var EaMusicIcon []byte

// 左侧导航栏
func (u *Ui) createLeftMenu() {
	// 推荐
	// title := canvas.Text{Text: "推荐"}
	// icon := fyne.NewStaticResource("EaMusicIcon", EaMusicIcon)
	// recommendBtn := widget.NewButtonWithIcon("精选音乐", icon, func() {
	//
	// })
	// u.FuncBar =
}
func (u *Ui) MakeUi() {

}

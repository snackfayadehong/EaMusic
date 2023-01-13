package ui

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Ui struct {
	MainWindow fyne.Window
	FuncBar    *fyne.Layout
}

var MyUi Ui

//go:embed EaMusic.png
var EaMusicIcon []byte

// 左侧导航栏
func (u *Ui) createLeftMenu() (*canvas.Text, *widget.Button) {
	// 推荐
	title := canvas.Text{Text: "推荐"}
	icon := fyne.NewStaticResource("EaMusicIcon", EaMusicIcon)
	recommendBtn := widget.NewButtonWithIcon("精选音乐", icon, func() {

	})
	return &title, recommendBtn
}
func (u *Ui) MakeUi() {
	title, recommend := u.createLeftMenu()
	funcBar := layout.NewGridLayout(2)
	u.FuncBar = &funcBar
}

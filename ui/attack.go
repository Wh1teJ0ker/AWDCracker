package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateAttackModule 创建攻击模块
func CreateAttackModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Attack Module"),
		widget.NewButton("Start Attack", func() {
			// 调用攻击功能
		}),
	)
}

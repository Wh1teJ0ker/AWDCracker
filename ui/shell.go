package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateShellModule 创建Shell模块
func CreateShellModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("Manage Shell", func() {
			// 管理shell文件
		}),
	)
}

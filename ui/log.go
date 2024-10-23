package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateLogModule 创建日志模块
func CreateLogModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Log Messages:"),
		// 这里可以添加显示日志内容的功能
		widget.NewLabel("No logs yet."),
	)
}

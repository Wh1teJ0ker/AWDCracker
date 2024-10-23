package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateScanModule 创建扫描模块
func CreateScanModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("Start Scan", func() {
			// 调用扫描功能
		}),
	)
}

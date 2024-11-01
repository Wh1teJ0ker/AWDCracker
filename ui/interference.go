package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateInterferenceModule 创建干扰模块
func CreateInterferenceModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("Start Interference", func() {
			// 调用干扰功能
		}),
	)
}

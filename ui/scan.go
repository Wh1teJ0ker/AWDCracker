package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateScanModule 创建扫描模块
func CreateScanModule() fyne.CanvasObject {
	// 初始化进度条
	progressBar := widget.NewProgressBar()
	progressBar.SetValue(0) // 设置初始值为 0

	// 创建扫描模块布局
	return container.NewVBox(
		container.NewGridWithColumns(3, // 设置为三列
			widget.NewLabel("IP Address:"),                     // 第一列：标签
			widget.NewSelectEntry([]string{"全IP探测", "单一IP探测"}), // 第二列：IP模式下拉列表
			widget.NewEntry(), // 第三列：IP 地址输入框

			widget.NewLabel("Port:"), // 第一列：标签
			widget.NewSelectEntry([]string{"全端口", "常用端口", "无端口设置"}), // 第二列：端口模式下拉列表
			widget.NewEntry(), // 第三列：端口输入框
		),
		container.NewHBox( // 按钮和进度条在同一行
			widget.NewButton("扫描", func() {
				// 在这里定义扫描功能
				progressBar.SetValue(0.5) // 示例：设置进度条的值为50%
			}),
			progressBar,
		),
	)
}

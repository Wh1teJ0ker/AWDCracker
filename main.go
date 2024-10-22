package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("AWDCracker")

	// 创建四个功能模块的界面
	scanModule := createScanModule()
	attackModule := createAttackModule()
	interferenceModule := createInterferenceModule()
	ShellModule := createShellModule()
	logModule := createLogModule()

	// 将模块并排放置
	w.SetContent(container.NewHBox(
		scanModule,
		attackModule,
		interferenceModule,
		ShellModule,
		logModule,
	))

	w.ShowAndRun()
}

// createScanModule 创建扫描模块
func createScanModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("Start Scan", func() {
			// 调用扫描功能
		}),
	)
}

// createAttackModule 创建攻击模块
func createAttackModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("Start Attack", func() {
			// 调用攻击功能
		}),
	)
}

// createInterferenceModule 创建干扰模块
func createInterferenceModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("Start Interference", func() {
			// 调用干扰功能
		}),
	)
}

func createShellModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("Manage Shell", func() {
			// 管理shell文件
		}),
	)
}

// createLogModule 创建日志模块
func createLogModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Log Messages:"),
		// 这里可以添加显示日志内容的功能
	)
}

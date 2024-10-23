package main

import (
	"AWDCracker/ui" // 导入 UI 包

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("AWDCracker")

	// 创建四个功能模块的界面
	// 如需继续添加功能模块，继续在此调用
	scanModule := ui.CreateScanModule()
	attackModule := ui.CreateAttackModule()
	interferenceModule := ui.CreateInterferenceModule()
	shellModule := ui.CreateShellModule()
	logModule := ui.CreateLogModule()

	// 将模块并排放置
	w.SetContent(container.NewHBox(
		scanModule,
		attackModule,
		interferenceModule,
		shellModule,
		logModule,
	))

	w.ShowAndRun()
}

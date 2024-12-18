package main

import (
	"AWDCracker/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("AWDCracker")

	// 创建功能模块的界面
	w.Resize(fyne.NewSize(720, 480))
	scanModule := src.CreateScanModule(w)                // 扫描模块
	attackModule := src.CreateAttackModule()             // 攻击模块
	interferenceModule := src.CreateInterferenceModule() // 干扰模块
	shellModule := src.CreateShellModule()               // Shell 模块

	// 将模块放置在选项卡中
	tabs := container.NewAppTabs(
		container.NewTabItem("扫描模块", scanModule),
		container.NewTabItem("攻击模块", attackModule),
		container.NewTabItem("干扰模块", interferenceModule),
		container.NewTabItem("Shell模块", shellModule),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)

	w.ShowAndRun()
}

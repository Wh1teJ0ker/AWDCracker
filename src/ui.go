package src

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateScanModule 构建扫描模块主界面
func CreateScanModule(window fyne.Window) fyne.CanvasObject {
	form := CreateScanForm(window)
	progressBar := CreateProgressBar()
	resultsList, scrollableResults := CreateResultsList()
	scanButton := CreateScanButton(progressBar, resultsList)

	return container.NewVBox(
		form,
		container.NewHBox(scanButton, progressBar),
		scrollableResults,
	)
}

// CreateScanForm 创建扫描表单
func CreateScanForm(window fyne.Window) fyne.CanvasObject {
	ipEntry := CreateIPEntry()
	ipModeSelect := CreateIPModeSelect(ipEntry)
	importButton := CreateImportButton(window)

	return container.NewVBox(
		container.NewGridWithColumns(2,
			widget.NewLabel("IP Mode:"),
			ipModeSelect,
		),
		container.NewGridWithColumns(2,
			ipEntry,
			importButton,
		),
	)
}

// CreateIPEntry 创建 IP 地址输入框
func CreateIPEntry() *widget.Entry {
	ipEntry := widget.NewEntry()
	ipEntry.Disable()
	ipEntry.SetPlaceHolder("请输入 IP 地址")
	return ipEntry
}

// CreateIPModeSelect 创建 IP 模式选择器
func CreateIPModeSelect(ipEntry *widget.Entry) *widget.Select {
	return widget.NewSelect([]string{"全网段嗅探", "单一IP探测", "导入IP列表"}, func(selected string) {
		switch selected {
		case "单一IP探测":
			ipEntry.SetText("192.168.1.1")
			ipEntry.Enable()
		case "全网段嗅探":
			ipEntry.SetText("192.168.1.0/24")
			ipEntry.Enable()
		case "导入IP列表":
			ipEntry.SetText("")
			ipEntry.Disable()
		default:
			ipEntry.SetText("")
			ipEntry.Disable()
		}
	})
}

// CreateImportButton 创建“导入 IP”按钮
func CreateImportButton(win fyne.Window) *widget.Button {
	return widget.NewButton("导入 IP", func() {

	})
}

// CreateProgressBar 创建进度条
func CreateProgressBar() *widget.ProgressBar {
	progressBar := widget.NewProgressBar()
	progressBar.SetValue(0)
	return progressBar
}

// CreateResultsList 创建结果列表
func CreateResultsList() (*widget.List, fyne.CanvasObject) {
	results := []string{}

	resultsList := widget.NewList(
		func() int {
			return len(results)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("扫描结果")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(results[i])
		},
	)

	scrollableResults := container.NewVScroll(resultsList)
	scrollableResults.SetMinSize(fyne.NewSize(600, 400))
	return resultsList, scrollableResults
}

// CreateScanButton 创建扫描按钮
func CreateScanButton(progressBar *widget.ProgressBar, resultsList *widget.List) *widget.Button {
	return widget.NewButton("扫描", func() {
		go simulateScan(progressBar)
	})
}

// simulateScan 模拟扫描进度
func simulateScan(progressBar *widget.ProgressBar) {
	for i := 0.0; i <= 1.0; i += 0.1 {
		time.Sleep(500 * time.Millisecond)
		progressBar.SetValue(i)
	}
}

func CreateAttackModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("Start Attack", func() {
			// 调用攻击功能
		}),
	)
}

func CreateInterferenceModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("Start Interference", func() {
			// 调用干扰功能
		}),
	)
}

func CreateShellModule() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewButton("Manage Shell", func() {
			// 管理shell文件
		}),
	)
}

package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type GUI struct {
	app    fyne.App
	window fyne.Window
}

func NewGUI() *GUI {
	return &GUI{
		app: app.New(),
	}
}

func (g *GUI) Start() {
	g.window = g.app.NewWindow("DeepSpace")

	// 加载 DeepSeek logo
	logo := canvas.NewImageFromFile("assets/1.jpg")
	logo.Resize(fyne.NewSize(200, 200))
	logo.FillMode = canvas.ImageFillOriginal

	// 创建主要控件
	startButton := widget.NewButton("启动服务", func() {
		go startServer("127.0.0.1", "9988")
	})

	portEntry := widget.NewEntry()
	portEntry.SetPlaceHolder("端口 (默认: 9988)")

	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("主机地址 (默认: 127.0.0.1)")

	// 创建请求列表
	requestList := widget.NewList(
		func() int { return 0 }, // 获取请求数量
		func() fyne.CanvasObject { // 创建列表项模板
			return widget.NewLabel("Request")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) { // 更新列表项
			label := item.(*widget.Label)
			label.SetText(fmt.Sprintf("Request %d", id))
		},
	)

	// 布局设置
	controls := container.NewVBox(
		hostEntry,
		portEntry,
		startButton,
	)

	content := container.NewHSplit(
		container.NewVBox(
			logo,
			controls,
		),
		requestList,
	)

	g.window.SetContent(content)
	g.window.Resize(fyne.NewSize(800, 600))
	g.window.CenterOnScreen()
	g.window.SetMaster()

	g.window.ShowAndRun()
}

func startServer(host, port string) {
	// 这里调用现有的服务器启动逻辑
	fmt.Printf("Starting server on %s:%s\n", host, port)
}

func StartGUI() {
	gui := NewGUI()
	gui.Start()
}

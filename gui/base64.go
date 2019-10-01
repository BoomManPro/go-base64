package main

import (
	"encoding/base64"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"go-base64/gui/windows"
	"io"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

func main() {

	width := 600
	height := 300
	config := windows.MainWindows{
		Initialized: true,
		InitWidth:   width,
		InitHeight:  height,
		Size: Size{
			Width:  width,
			Height: height,
		},
		MinSize: Size{
			Width:  width,
			Height: height,
		},
	}

	window := initMainWindows(&config)

	if _, e := window.Run(); e != nil {
		panic(e)
	}

}

func handlerFile(reader io.Reader) string {
	bytes, e := ioutil.ReadAll(reader)
	if e != nil {
		return e.Error()
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func initMainWindows(config *windows.MainWindows) MainWindow {
	var textEdit *walk.TextEdit
	var mw *walk.MainWindow
	var content string

	window := MainWindow{
		AssignTo: &mw,
		Title:    "Base64Encoding",
		MinSize:  config.MinSize,

		Layout: VBox{},
		Size:   config.Size,
		OnDropFiles: func(files []string) {
			var viewContent string

			file, e := os.Open(files[0])
			if e != nil {
				viewContent = e.Error()
			} else {
				viewContent = handlerFile(file)
				content = viewContent
			}
			textEdit.SetText(getSomeViewContent(viewContent))
		},
		Children: []Widget{
			PushButton{
				Text: "Copy",
				OnClicked: func() {
					if err := walk.Clipboard().SetText(content); err != nil {
						log.Print("Copy: ", err)
					}
				},
			},
			TextEdit{
				AssignTo: &textEdit,
				ReadOnly: true,
				Text:     "Drop files here, from windows explorer...,file encoding base64 and clip copy",
			},
		},
		OnBoundsChanged: func() {
			if config.Initialized {
				//println(`宽度：`, GetSystemMetrics(0))
				//println(`高度：`, GetSystemMetrics(1))

				mw.SetBounds(walk.Rectangle{
					X:      (GetSystemMetrics(0) - config.InitWidth) / 2,
					Y:      (GetSystemMetrics(1) - config.InitHeight) / 2,
					Width:  config.InitHeight,
					Height: config.InitHeight,
				})
				config.Initialized = false
			}
		},
	}

	return window
}

func GetSystemMetrics(nIndex int) int {
	ret, _, _ := syscall.NewLazyDLL(`User32.dll`).NewProc(`GetSystemMetrics`).Call(uintptr(nIndex))
	return int(ret)
}

func getSomeViewContent(content string) string {
	if len(content) > 1000 {
		return content[0:1000] + "......"
	}
	return content

}

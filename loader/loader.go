package main

import (
	"os/exec"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"github.com/tink-ab/tempfile"
)

const (
	SIZE_W = 300
	SIZE_H = 200
)

func main() {
	var mw *walk.MainWindow

	MainWindow{
		Title:    "Demo",
		AssignTo: &mw,
		MinSize:  Size{300, 200},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:          "Loading ...",
				TextAlignment: AlignCenter,
				Font: Font{
					Family:    "Segoe UI",
					PointSize: 18,
					Bold:      true,
				},
			},
		},
	}.Create()

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER)

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)

	go say()
	mw.Run()
}

func say() {
	box := packr.NewBox("./content")
	app := box.Bytes("app.exe")
	tmpFile, _ := tempfile.TempFile("", "launc-", ".exe")
	tmpFile.Write(app)
	tmpFile.Close()

	time.Sleep(300)

	c := exec.Command(tmpFile.Name())
	if err := c.Run(); err != nil {
	}
}

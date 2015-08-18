// +build windows

package goui

import (
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

var mainWindow *walk.MainWindow

var windowStore = make(map[int]walk.Form)

func runWindowsWebView() {
	var le *walk.LineEdit
	var wv *walk.WebView

	fmt.Println("Run MainWindow!")
	MainWindow{
		AssignTo: &mainWindow,
		Title:    "Walk WebView Example",
		MinSize:  Size{800, 600},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			LineEdit{
				AssignTo: &le,
				Text:     Bind("wv.URL"),
				OnKeyDown: func(key walk.Key) {
					if key == walk.KeyReturn {
						wv.SetURL(le.Text())
					}
				},
			},
			WebView{
				AssignTo: &wv,
				Name:     "wv",
				URL:      "http://golang.org",
			},
		},
	}.Run()
}

func osInit() {
	// runtime.LockOSThread()
	guiReadyCallback()
	runWindowsWebView()
	// fmt.Scanln()
	// C.StartApp()
}

func osOpenWindow(window *Window, url string, styleFlags int) {

	fmt.Println("osOpenWindow: ", window, url, styleFlags)

	var dia *walk.Dialog
	var le *walk.LineEdit
	var wv *walk.WebView

	windowStore[window.handle] = dia

	Dialog{
		AssignTo: &dia,
		Title:    "Window Title",
		MinSize:  Size{800, 600},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			LineEdit{
				AssignTo: &le,
				Text:     Bind("wv.URL"),
				OnKeyDown: func(key walk.Key) {
					if key == walk.KeyReturn {
						wv.SetURL(le.Text())
					}
				},
			},
			WebView{
				AssignTo: &wv,
				Name:     "wv",
				URL:      url,
			},
		},
	}.Run(mainWindow)
}

func osStop() {
	// C.StopApp()
	// mainWindow.Close()
}

func osCloseWindow(window *Window) {
	if w, ok := windowStore[window.handle]; ok {
		w.AsFormBase().Close()
		delete(windowStore, window.handle)
	}
}

func osSetWindowTitle(window *Window, title string) {
	if w, ok := windowStore[window.handle]; ok {
		w.AsFormBase().SetTitle(title)
	}
}

func osGetScreenSize() (int, int) {
	w := win.GetSystemMetrics(win.SM_CXSCREEN)
	h := win.GetSystemMetrics(win.SM_CYSCREEN)
	return int(w), int(h)
}

func osSetWindowSize(window *Window, width int, height int) {
	if w, ok := windowStore[window.handle]; ok {
		w.SetSize(walk.Size{width, height})
	}
}

func osSetWindowPosition(window *Window, left int, top int) {
	if w, ok := windowStore[window.handle]; ok {
		w.SetX(left)
		w.SetY(top)
	}
}

func osRememberGeometry(window *Window, key string) {
	// C.RememberWindowGeometry(C.int(window.handle), C.CString(key))
}

func osRunModal(window *Window) {
	// if w, ok := windowStore[window.handle]; ok {
	// 	w.Run()
	// }
}

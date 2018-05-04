package ui

import (
	"github.com/therecipe/qt/widgets"
)

type MainWindow struct {
	*widgets.QMainWindow
}

func NewMainWindow() *MainWindow {
	window := &MainWindow{widgets.NewQMainWindow(nil, 0)}
	window.init()
	return window
}

func (m *MainWindow) init() {
	frame := NewMainFrame(m, 0)
	m.SetCentralWidget(frame)
}

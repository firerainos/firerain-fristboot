package main

import (
	"github.com/firerainos/firerain-fristboot/ui"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	mainwindow := ui.NewMainWindow()
	mainwindow.Show()

	os.Exit(widgets.QApplication_Exec())
}

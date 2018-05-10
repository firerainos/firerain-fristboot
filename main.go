package main

import (
	"github.com/firerainos/firerain-fristboot/ui"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	app :=widgets.NewQApplication(len(os.Args), os.Args)
	app.SetApplicationName("firerain-fristboot")
	app.SetApplicationVersion("0.0.1")

	mainwindow := ui.NewMainWindow()
	//mainwindow.Show()
	mainwindow.SetFixedSize(app.Desktop().Geometry().Size())
	mainwindow.ShowFullScreen()

	os.Exit(app.Exec())
}

package page

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type EndPage struct {
	*widgets.QFrame
}

func NewEndPage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *EndPage {
	page := &EndPage{widgets.NewQFrame(parent, fo)}

	page.init()

	return page
}

func (page *EndPage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(page)
	vboxLayout.SetContentsMargins(0, 0, 0, 0)

	label := widgets.NewQLabel2("请稍等...",page,0)

	vboxLayout.AddStretch(1)
	vboxLayout.AddWidget(label,0,core.Qt__AlignCenter)
	vboxLayout.AddStretch(1)

	page.SetLayout(vboxLayout)
}

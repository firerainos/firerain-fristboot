package page

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type EndPage struct {
	*widgets.QFrame

	tipsLabel *widgets.QLabel
}

func NewEndPage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *EndPage {
	page := &EndPage{QFrame:widgets.NewQFrame(parent, fo)}

	page.init()

	return page
}

func (page *EndPage) init() {
	vboxLayout := widgets.NewQVBoxLayout2(page)
	vboxLayout.SetContentsMargins(0, 0, 0, 0)

	page.tipsLabel = widgets.NewQLabel2("请稍等...",page,0)

	vboxLayout.AddStretch(1)
	vboxLayout.AddWidget(page.tipsLabel,0,core.Qt__AlignCenter)
	vboxLayout.AddStretch(1)

	page.SetLayout(vboxLayout)
}

func (page *EndPage) AddTips(tips string) {
	page.tipsLabel.SetText(page.tipsLabel.Text()+"\n"+tips)
}

package ui

import (
	"github.com/firerainos/firerain-fristboot/ui/page"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type MainFrame struct {
	*widgets.QFrame

	welcomePage *page.WelcomePage
	accountPage *page.AccountPage
	endPage *page.EndPage

	backButton, nextButton *widgets.QPushButton

	stackLayout *widgets.QStackedLayout
}

func NewMainFrame(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *MainFrame {
	frame := &MainFrame{QFrame: widgets.NewQFrame(parent, fo)}

	frame.init()
	frame.initConnect()

	return frame
}

func (m *MainFrame) init() {
	vboxLayout := widgets.NewQVBoxLayout2(m)
	vboxLayout.SetContentsMargins(0, 0, 0, 0)

	m.stackLayout = widgets.NewQStackedLayout()

	m.welcomePage = page.NewWelcomePage(m, 0)
	m.accountPage = page.NewAccountPage(m, 0)
	m.endPage = page.NewEndPage(m, 0)

	m.backButton = widgets.NewQPushButton2("返回", m)
	m.nextButton = widgets.NewQPushButton2("继续", m)

	m.backButton.SetVisible(false)

	hboxLayout := widgets.NewQHBoxLayout()

	hboxLayout.AddStretch(1)
	hboxLayout.AddWidget(m.backButton, 0, core.Qt__AlignHCenter)
	hboxLayout.AddWidget(m.nextButton, 0, core.Qt__AlignHCenter)
	hboxLayout.AddStretch(1)

	m.stackLayout.AddWidget(m.welcomePage)
	m.stackLayout.AddWidget(m.accountPage)
	m.stackLayout.AddWidget(m.endPage)

	vboxLayout.AddLayout(m.stackLayout, 1)
	vboxLayout.AddLayout(hboxLayout, 1)
	vboxLayout.AddSpacing(20)

	m.SetLayout(vboxLayout)
}

func (m *MainFrame) initConnect() {
	m.stackLayout.ConnectCurrentChanged(func(index int) {
		if index == 0 {
			m.backButton.SetVisible(false)
		} else if index == m.stackLayout.Count()-1 {
			m.nextButton.SetVisible(false)
			m.backButton.SetVisible(false)
		} else if index > 0 {
			m.backButton.SetVisible(true)
		}
	})

	m.backButton.ConnectClicked(func(checked bool) {
		if m.nextButton.Text() == "确定"{
			m.nextButton.SetText("继续")
		}
		m.stackLayout.SetCurrentIndex(m.stackLayout.CurrentIndex() - 1)
	})

	m.nextButton.ConnectClicked(func(checked bool) {
		switch m.nextButton.Text() {
		case "继续":
			m.nextButton.SetText("确定")
		case "确定":
			password := m.accountPage.Password.Text()
			againPassword := m.accountPage.AgainPassword.Text()
			if password != againPassword {
				m.accountPage.SetTips("确认密码不匹配")
				return
			}
		}
		m.stackLayout.SetCurrentIndex(m.stackLayout.CurrentIndex() + 1)
	})
}
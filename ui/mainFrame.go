package ui

import (
	core2 "github.com/firerainos/firerain-fristboot/core"
	"github.com/firerainos/firerain-fristboot/styles"
	"github.com/firerainos/firerain-fristboot/ui/page"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
)

type MainFrame struct {
	*widgets.QFrame

	welcomePage *page.WelcomePage
	accountPage *page.AccountPage
	endPage     *page.EndPage

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

	m.backButton.SetMinimumWidth(60)
	m.nextButton.SetMinimumWidth(60)

	m.backButton.SetStyleSheet(styles.BackButton)
	m.nextButton.SetStyleSheet(styles.NextButton)

	m.backButton.SetVisible(false)

	hboxLayout := widgets.NewQHBoxLayout()
	hboxLayout.SetSpacing(40)

	hboxLayout.AddStretch(1)
	hboxLayout.AddWidget(m.backButton, 0, core.Qt__AlignHCenter)
	hboxLayout.AddWidget(m.nextButton, 0, core.Qt__AlignHCenter)
	hboxLayout.AddStretch(1)

	m.stackLayout.AddWidget(m.welcomePage)
	m.stackLayout.AddWidget(m.accountPage)
	m.stackLayout.AddWidget(m.endPage)

	vboxLayout.AddLayout(m.stackLayout, 1)
	vboxLayout.AddSpacing(50)
	vboxLayout.AddLayout(hboxLayout, 1)
	vboxLayout.AddSpacing(50)

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
		if m.nextButton.Text() == "确定" {
			m.nextButton.SetText("继续")
		}
		m.stackLayout.SetCurrentIndex(m.stackLayout.CurrentIndex() - 1)
	})

	m.nextButton.ConnectClicked(func(checked bool) {
		switch m.nextButton.Text() {
		case "继续":
			m.nextButton.SetText("确定")
		case "确定":
			if !m.accountPage.Check() {
				return
			}
			go m.configuration()
		case "开始使用":
			os.Exit(0)
		}
		m.stackLayout.SetCurrentIndex(m.stackLayout.CurrentIndex() + 1)
	})
}

func (m *MainFrame) configuration() {
	username := m.accountPage.Username.Text()

	if err := core2.UserAdd(username, m.accountPage.Password.Text()); err != nil {
		m.endPage.AddTips("用户创建失败,请在点击开始使用后切换到tty2手动创建:\n")
	}

	if err := core2.SetHomeName(m.accountPage.Hostname.Text()); err != nil {
		m.endPage.AddTips("主机名设置失败,可手动设置:\n")
	}

	if err := core2.SetLocale(); err != nil {
		m.endPage.AddTips("Locale设置失败,可手动设置:\n" + err.Error())
	}

	if err := core2.SetIM(); err != nil {
		m.endPage.AddTips("输入法环境变量设置失败,可手动设置:\n" + err.Error())
	}

	//if !core2.RemoveFristBoot() {
	//	m.endPage.AddTips("FristBoot卸载失败，请在登陆后手动卸载(sudo pacman -Rscn firerain-fristboot)")
	//}

	m.nextButton.SetText("开始使用")
	m.nextButton.SetVisible(true)
}

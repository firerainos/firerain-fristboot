package page

import (
	widgets2 "github.com/firerainos/firerain-fristboot/ui/widgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"log"
	"os"
)

type AccountPage struct {
	*widgets.QFrame

	tipsLabel *widgets.QLabel

	Username, Hostname, Password, AgainPassword *widgets2.LineEdit
}

func NewAccountPage(parent widgets.QWidget_ITF, fo core.Qt__WindowType) *AccountPage {
	frame := widgets.NewQFrame(parent, fo)

	accountPage := &AccountPage{QFrame: frame}
	accountPage.init()
	accountPage.initConnect()

	return accountPage
}

func (page *AccountPage) init() {
	pwd, _ := os.Getwd()
	log.Println(pwd)

	vboxLayout := widgets.NewQVBoxLayout2(page)

	logoLabel := widgets.NewQLabel(page, 0)

	logoLabel.SetPixmap(gui.NewQPixmap5(pwd+"/resources/logo.png", "", 0).
		Scaled2(150, 150, core.Qt__KeepAspectRatioByExpanding, 0))
	logoLabel.SetFixedSize2(150, 150)

	titleLabel := widgets.NewQLabel2("创建用户帐号", page, 0)

	page.tipsLabel = widgets.NewQLabel2("输入用户名和密码", page, 0)
	page.tipsLabel.SetFixedWidth(200)
	page.tipsLabel.SetAlignment(core.Qt__AlignCenter)

	page.Username = widgets2.NewLineEdit(pwd+"/resources/username.svg", page)
	page.Hostname = widgets2.NewLineEdit(pwd+"/resources/host.svg", page)
	page.Password = widgets2.NewLineEdit(pwd+"/resources/password.svg", page)
	page.AgainPassword = widgets2.NewLineEdit(pwd+"/resources/password.svg", page)

	regExp := core.NewQRegExp()
	regExp.SetPattern("^[a-z][-a-z0-9_]*$")

	page.Username.SetValidator(gui.NewQRegExpValidator2(regExp, page))

	page.Username.SetPlaceholderText("用户名")
	page.Hostname.SetPlaceholderText("主机名")
	page.Password.SetPlaceholderText("用户密码")
	page.AgainPassword.SetPlaceholderText("确认用户密码")

	page.Password.SetEchoMode(widgets.QLineEdit__Password)
	page.AgainPassword.SetEchoMode(widgets.QLineEdit__Password)

	vboxLayout.AddSpacing(100)
	vboxLayout.AddWidget(logoLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddSpacing(30)
	vboxLayout.AddWidget(titleLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddWidget(page.tipsLabel, 0, core.Qt__AlignHCenter)
	vboxLayout.AddSpacing(100)
	vboxLayout.AddStretch(10)
	vboxLayout.AddWidget(page.Username, 0, core.Qt__AlignCenter)
	vboxLayout.AddWidget(page.Hostname, 0, core.Qt__AlignCenter)
	vboxLayout.AddWidget(page.Password, 0, core.Qt__AlignCenter)
	vboxLayout.AddWidget(page.AgainPassword, 0, core.Qt__AlignCenter)
	vboxLayout.AddStretch(1)

	page.SetLayout(vboxLayout)
}
func (page *AccountPage) initConnect() {
	page.Username.ConnectTextChanged(func(text string) {
		page.Hostname.SetText(text + "-PC")
	})
}

func (page *AccountPage) SetTips(tips string) {
	page.tipsLabel.SetText(tips)
}

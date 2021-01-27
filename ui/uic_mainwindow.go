package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type __mainwindow struct{}

func (*__mainwindow) init() {}

type MainWindow struct {
	*__mainwindow
	*widgets.QMainWindow
	Centralwidget *widgets.QWidget
	SendToButton  *widgets.QPushButton
	TextEdit      *widgets.QTextEdit
	KafkaButton   *widgets.QPushButton
	IpLineEdit    *widgets.QLineEdit
	IpLable       *widgets.QLabel
	PortLable     *widgets.QLabel
	PortLineEdit  *widgets.QLineEdit
	Label         *widgets.QLabel
	TopicLineEdit *widgets.QLineEdit
	Menubar       *widgets.QMenuBar
	Statusbar     *widgets.QStatusBar
}

func NewMainWindow(p widgets.QWidget_ITF) *MainWindow {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	file := core.NewQFile2(":/ui/mainwindow.ui")
	file.Open(core.QIODevice__ReadOnly)
	w := &MainWindow{QMainWindow: widgets.NewQMainWindowFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}
	file.Close()
	w.setupUI()
	w.init()
	return w
}
func (w *MainWindow) setupUI() {
	w.Centralwidget = widgets.NewQWidgetFromPointer(w.FindChild("centralwidget", core.Qt__FindChildrenRecursively).Pointer())
	w.SendToButton = widgets.NewQPushButtonFromPointer(w.FindChild("sendToButton", core.Qt__FindChildrenRecursively).Pointer())
	w.KafkaButton = widgets.NewQPushButtonFromPointer(w.FindChild("kafkaButton", core.Qt__FindChildrenRecursively).Pointer())
	w.IpLineEdit = widgets.NewQLineEditFromPointer(w.FindChild("ipLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.IpLable = widgets.NewQLabelFromPointer(w.FindChild("ipLable", core.Qt__FindChildrenRecursively).Pointer())
	w.Label = widgets.NewQLabelFromPointer(w.FindChild("label", core.Qt__FindChildrenRecursively).Pointer())
	w.Statusbar = widgets.NewQStatusBarFromPointer(w.FindChild("statusbar", core.Qt__FindChildrenRecursively).Pointer())
	w.TextEdit = widgets.NewQTextEditFromPointer(w.FindChild("textEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.PortLable = widgets.NewQLabelFromPointer(w.FindChild("portLable", core.Qt__FindChildrenRecursively).Pointer())
	w.PortLineEdit = widgets.NewQLineEditFromPointer(w.FindChild("portLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.TopicLineEdit = widgets.NewQLineEditFromPointer(w.FindChild("topicLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.Menubar = widgets.NewQMenuBarFromPointer(w.FindChild("menubar", core.Qt__FindChildrenRecursively).Pointer())
}

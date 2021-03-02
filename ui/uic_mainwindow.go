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
	Centralwidget              *widgets.QWidget
	TextEdit                   *widgets.QTextEdit
	HorizontalLayoutWidget     *widgets.QWidget
	HorizontalLayout_kafkaInfo *widgets.QHBoxLayout
	IpLable                    *widgets.QLabel
	IpLineEdit                 *widgets.QLineEdit
	PortLable                  *widgets.QLabel
	PortLineEdit               *widgets.QLineEdit
	Label                      *widgets.QLabel
	TopicLineEdit              *widgets.QLineEdit
	KafkaButton                *widgets.QPushButton
	VerticalLayoutWidget       *widgets.QWidget
	VerticalLayout_send        *widgets.QVBoxLayout
	SendLogBrowser             *widgets.QTextBrowser
	SendToButton               *widgets.QPushButton
	VerticalLayoutWidget_2     *widgets.QWidget
	VerticalLayout_radio       *widgets.QVBoxLayout
	GroupBox_textType          *widgets.QGroupBox
	RadioButton_Json           *widgets.QRadioButton
	RadioButton_Raw            *widgets.QRadioButton
	Menubar                    *widgets.QMenuBar
	Statusbar                  *widgets.QStatusBar
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
	w.IpLineEdit = widgets.NewQLineEditFromPointer(w.FindChild("ipLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.PortLable = widgets.NewQLabelFromPointer(w.FindChild("portLable", core.Qt__FindChildrenRecursively).Pointer())
	w.SendToButton = widgets.NewQPushButtonFromPointer(w.FindChild("sendToButton", core.Qt__FindChildrenRecursively).Pointer())
	w.RadioButton_Raw = widgets.NewQRadioButtonFromPointer(w.FindChild("radioButton_Raw", core.Qt__FindChildrenRecursively).Pointer())
	w.Menubar = widgets.NewQMenuBarFromPointer(w.FindChild("menubar", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayoutWidget = widgets.NewQWidgetFromPointer(w.FindChild("horizontalLayoutWidget", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_kafkaInfo = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_kafkaInfo", core.Qt__FindChildrenRecursively).Pointer())
	w.IpLable = widgets.NewQLabelFromPointer(w.FindChild("ipLable", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout_radio = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout_radio", core.Qt__FindChildrenRecursively).Pointer())
	w.RadioButton_Json = widgets.NewQRadioButtonFromPointer(w.FindChild("radioButton_Json", core.Qt__FindChildrenRecursively).Pointer())
	w.Statusbar = widgets.NewQStatusBarFromPointer(w.FindChild("statusbar", core.Qt__FindChildrenRecursively).Pointer())
	w.TextEdit = widgets.NewQTextEditFromPointer(w.FindChild("textEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.Label = widgets.NewQLabelFromPointer(w.FindChild("label", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayoutWidget = widgets.NewQWidgetFromPointer(w.FindChild("verticalLayoutWidget", core.Qt__FindChildrenRecursively).Pointer())
	w.SendLogBrowser = widgets.NewQTextBrowserFromPointer(w.FindChild("sendLogBrowser", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayoutWidget_2 = widgets.NewQWidgetFromPointer(w.FindChild("verticalLayoutWidget_2", core.Qt__FindChildrenRecursively).Pointer())
	w.GroupBox_textType = widgets.NewQGroupBoxFromPointer(w.FindChild("groupBox_textType", core.Qt__FindChildrenRecursively).Pointer())
	w.Centralwidget = widgets.NewQWidgetFromPointer(w.FindChild("centralwidget", core.Qt__FindChildrenRecursively).Pointer())
	w.PortLineEdit = widgets.NewQLineEditFromPointer(w.FindChild("portLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.TopicLineEdit = widgets.NewQLineEditFromPointer(w.FindChild("topicLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.KafkaButton = widgets.NewQPushButtonFromPointer(w.FindChild("kafkaButton", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout_send = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout_send", core.Qt__FindChildrenRecursively).Pointer())
}

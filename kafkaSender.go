package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"kafkaSender/ui"
	"os"
)

func main(){
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling,true)
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window:=ui.NewMainWindow(nil)

	var kafkaConnect *kafka.Conn
	var err error

	//window.TextEdit.TextChanged()
	window.RadioButton_Json.ConnectClicked(func(checked bool) {
		fmt.Println("checked")
		if checked{
			
		}})
	window.KafkaButton.ConnectClicked(func(checked bool) {
		fmt.Println("come here...")
		if len(window.IpLineEdit.Text())==0{
			widgets.QMessageBox_Information(nil,"Ok","Please input ip address!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		if len(window.PortLineEdit.Text())==0{
			widgets.QMessageBox_Information(nil,"Ok","Please input port!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		if len(window.TopicLineEdit.Text())==0{
			widgets.QMessageBox_Information(nil,"Ok","Please input topic", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		if kafkaConnect!=nil{
			err:=kafkaConnect.Close()
			if err!=nil{
				fmt.Println("close err",err.Error())
			}
		}
		window.SendLogBrowser.SetText(window.SendLogBrowser.ToPlainText()+"connct to "+window.IpLineEdit.Text()+":"+window.PortLineEdit.Text()+"\n")
		kafkaConnect, err = kafka.DialLeader(context.Background(), "tcp", window.IpLineEdit.Text()+":"+window.PortLineEdit.Text(), window.TopicLineEdit.Text(), 0)

		if err!=nil{
			widgets.QMessageBox_Information(nil,"Ok",err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}else {
			widgets.QMessageBox_Information(nil,"Ok","Connect success!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
	})

	window.SendToButton.ConnectClicked(func(checked bool) {
		if kafkaConnect==nil {
			widgets.QMessageBox_Information(nil,"ERROR","Please connect kafka first!",widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		fmt.Println(window.TextEdit.ToPlainText())
		_,err:=kafkaConnect.Write([]byte(window.TextEdit.ToPlainText()))
		if err!=nil{
			widgets.QMessageBox_Information(nil,"ERROR",err.Error()+"\nPlease reconnect kafka！",widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}else{
			window.SendLogBrowser.SetText(window.SendLogBrowser.ToPlainText()+"send success!\n")
		}
	})

	defer func() {
		if kafkaConnect!=nil{
			fmt.Println("kafka close")
			err:=kafkaConnect.Close()
			if err!=nil{
				fmt.Println(err.Error())
			}
		}
	}()

	// make the window visible
	window.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	app.Exec()
}

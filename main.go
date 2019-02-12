package main

import (
	"io/ioutil"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	iconData, err := getIconBytes("ServiceNowImg.ico")
	if err == nil {
		systray.SetIcon(iconData)
	}
	systray.SetTitle("Servive Now Alerts")
	systray.SetTooltip("Service Now Alerts")

	setupExitMenu()

}

func setupExitMenu() {
	//setup exit menu
	mQuitMenuItem := systray.AddMenuItem("Quit", "Quit the app")
	quitIconData, err := getIconBytes("off.ico")
	if err == nil {
		mQuitMenuItem.SetIcon(quitIconData)
	}
	//setup exit go routine
	go func() {
		<-mQuitMenuItem.ClickedCh
		systray.Quit()
	}()
}

func onExit() {

}

func getIconBytes(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

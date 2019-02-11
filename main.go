package main

import (
	"io/ioutil"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	iconData, err := getIconBytes()
	if err == nil {
		systray.SetIcon(iconData)
	}
	systray.SetTitle("Servive Now Alerts")
	systray.SetTooltip("The tips of spear")
	mQuit := systray.AddMenuItem("Quit", "Quit the app")

}

func onExit() {

}

func exit(mQuitMenuItem *MenuItem) {

}

func getIconBytes() ([]byte, error) {
	return ioutil.ReadFile("ServiceNowImg.ico")
}

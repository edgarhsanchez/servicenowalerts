package main

import (
	"io/ioutil"

	"github.com/getlantern/systray"
	"github.com/zserge/webview"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	iconData, err := getIconBytes("resources/ServiceNowImg.ico")
	if err == nil {
		systray.SetIcon(iconData)
	}
	systray.SetTitle("Servive Now Alerts")
	systray.SetTooltip("Service Now Alerts")

	setupCredsMenu()

	setupExitMenu()

}

func setupExitMenu() {
	//setup exit menu
	systray.AddSeparator()
	mQuitMenuItem := systray.AddMenuItem("Quit", "Quit the app")
	quitIconData, err := getIconBytes("resources/off.ico")
	if err == nil {
		mQuitMenuItem.SetIcon(quitIconData)
	}
	//setup exit go routine
	go func() {
		<-mQuitMenuItem.ClickedCh
		systray.Quit()
	}()
}

func setupCredsMenu() {
	//setup credentials launching menu item
	mLaunchCredsMenuItem := systray.AddMenuItem("Enter Credentials", "Launch Window to Enter Service Now Credentials")
	launchCredsIconData, err := getIconBytes("resources/Key.ico")
	if err == nil {
		mLaunchCredsMenuItem.SetIcon(launchCredsIconData)
	}

	// setup launch creds menu item go routing
	go openWebView(mLaunchCredsMenuItem)
}

func onExit() {

}

func getIconBytes(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

func openWebView(mLaunchCredsMenuItem *systray.MenuItem) {
	<-mLaunchCredsMenuItem.ClickedCh
	webview.Open("Service Now Alerts", "https://www.google.com", 800, 600, true)
	go openWebView(mLaunchCredsMenuItem)
}

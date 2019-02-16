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

	webView := createWebView()

	setupCredsMenu(webView)

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

func setupCredsMenu(webView *webview.WebView) {
	//setup credentials launching menu item
	mLaunchCredsMenuItem := systray.AddMenuItem("Enter Credentials", "Launch Window to Enter Service Now Credentials")
	launchCredsIconData, err := getIconBytes("resources/Key.ico")
	if err == nil {
		mLaunchCredsMenuItem.SetIcon(launchCredsIconData)
	}

	// setup launch creds menu item go routing
	go setupOpenCredsEvent(webView, mLaunchCredsMenuItem)
}

func onExit() {

}

func getIconBytes(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

func createWebView() *webview.WebView {
	webView := webview.New(webview.Settings{
		Title:     "Service Now Alerts",
		URL:       "https://www.google.com",
		Width:     800,
		Height:    600,
		Resizable: true,
	})
	defer webView.Exit()
	webView.SetFullscreen(true)
	webView.Run()
	return &webView
}

func setupOpenCredsEvent(webView *webview.WebView, mLaunchCredsMenuItem *systray.MenuItem) {
	<-mLaunchCredsMenuItem.ClickedCh
	(*webView).Run()
	go setupOpenCredsEvent(webView, mLaunchCredsMenuItem)
}

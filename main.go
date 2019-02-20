package main

import (
	"io/ioutil"
	"net/http"

	"github.com/awnumar/memguard"

	"github.com/getlantern/systray"
	"github.com/zserge/webview"
)

type ServiceNow struct {
	url   string
	title string
}

func (svcNow *ServiceNow) saveServiceNowInfo(serviceNowUrl string, userName string, password string) {
	_userName = userName
	_password, _ = memguard.NewMutableFromBytes([]byte(password))
}

var (
	_secret_key string
	_viewUrl    string
	_userName   string
	_password   *memguard.LockedBuffer
	_serviceNow ServiceNow
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

	setupWebPage()
	webView := createWebView()

	setupCredsMenu(webView)

	setupExitMenu()

}

func setupWebPage() {
	go func() {
		http.Handle("/", http.FileServer(assetFS()))
		err := http.ListenAndServe(":9999", nil)
		if err != nil {
			panic("http listening caused a panic")
		}
	}()
	_viewUrl = "http://127.0.0.1:9999/index.html"

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
		Title:  "Service Now Alerts",
		URL:    _viewUrl,
		Width:  800,
		Height: 600,
	})
	defer webView.Exit()
	webView.SetFullscreen(true)
	webView.Dispatch(func() {
		_serviceNow.url = "https://www.google.com"
		webView.Bind("svc_now", &_serviceNow)
	})
	webView.Run()
	return &webView
}

func setupOpenCredsEvent(webView *webview.WebView, mLaunchCredsMenuItem *systray.MenuItem) {
	<-mLaunchCredsMenuItem.ClickedCh
	(*webView).Exit()
	webView = createWebView()
	go setupOpenCredsEvent(webView, mLaunchCredsMenuItem)
}

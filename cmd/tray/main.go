package main

import "github.com/getlantern/systray"

func main() {
	onExit := func() {
		// now := time.Now()
		// ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("BDN9")
	systray.SetTooltip("Lantern")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the app")

	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()
}

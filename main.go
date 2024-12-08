package main

import "github.com/go-rod/rod"

func main() {
	page := rod.New().MustConnect().MustPage("https://google.com")
	page.MustWaitStable().MustScreenshot("a.png")
}

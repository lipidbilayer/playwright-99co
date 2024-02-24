package home

import (
	"fmt"

	pw "github.com/playwright-community/playwright-go"
)

type HomeElement struct {
	page pw.Page
	url  string
}

func New(page pw.Page) *HomeElement {
	home := HomeElement{page: page, url: "https://www.99.co/id"}
	home.OpenPage()
	return &home
}

func (home *HomeElement) OpenPage() {
	_, err := home.page.Goto(home.url)
	if err != nil {
		panic(fmt.Sprintf("failed to open page with err %s", err.Error()))
	}
	err = home.page.WaitForLoadState(pw.PageWaitForLoadStateOptions{State: pw.LoadStateNetworkidle})
	// err = response.Finished()
	if err != nil {
		panic(fmt.Sprintf("failed to open page with err %s", err.Error()))
	}
}

func (home *HomeElement) GetPageBody() pw.Locator {
	return home.page.Locator("//*[@id=\"__next\"]")
}

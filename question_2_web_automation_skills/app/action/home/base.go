package home

import (
	he "99.co/web-test/app/pages_element/home"
	pw "github.com/playwright-community/playwright-go"
)

var (
	timeout = 5000.0
)

type HomeAction struct {
	element     *he.HomeElement
	browserPage pw.Page
	assertion   pw.PlaywrightAssertions
}

func New(browserPage pw.Page) *HomeAction {
	assertion := pw.NewPlaywrightAssertions(5000)
	homeElement := he.New(browserPage)
	action := HomeAction{element: homeElement, browserPage: browserPage, assertion: assertion}
	return &action
}

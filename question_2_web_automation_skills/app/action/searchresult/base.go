package searchresult

import (
	el "99.co/web-test/app/pages_element/searchresult"
	pw "github.com/playwright-community/playwright-go"
)

var (
	timeout = 5000.0
)

type SearchResultAction struct {
	element     *el.SearchElement
	browserPage pw.Page
	assertion   pw.PlaywrightAssertions
}

func New(browserPage pw.Page) *SearchResultAction {
	assertion := pw.NewPlaywrightAssertions(timeout)
	homeElement := el.New(browserPage)
	action := SearchResultAction{element: homeElement, browserPage: browserPage, assertion: assertion}
	return &action
}

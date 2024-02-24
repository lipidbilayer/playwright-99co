package searchresult

import (
	pw "github.com/playwright-community/playwright-go"
)

type SearchElement struct {
	page pw.Page
	url  string
}

func New(page pw.Page) *SearchElement {
	search := SearchElement{page: page, url: page.URL()}

	return &search
}

func (elem *SearchElement) GetPageBody() pw.Locator {
	return elem.page.Locator("//*[@id=\"__next\"]")
}

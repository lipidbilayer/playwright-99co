package searchresult

import (
	pw "github.com/playwright-community/playwright-go"
)

func (elem *SearchElement) GetSearchResultElementItems() ([]pw.Locator, error) {
	return elem.page.Locator("//div[@class='srpListing']/div[@class='cardSecondary']").All()
}

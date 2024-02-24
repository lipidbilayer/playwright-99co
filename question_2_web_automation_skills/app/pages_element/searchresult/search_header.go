package searchresult

import pw "github.com/playwright-community/playwright-go"

func (elem *SearchElement) GetSearchResultElementHeaderTitle() pw.Locator {
	return elem.page.Locator("//h1")
}

func (elem *SearchElement) GetSearchResultElementHeaderBreadcrumb() pw.Locator {
	return elem.page.Locator("//div[@class='breadcrumbSEO']")
}

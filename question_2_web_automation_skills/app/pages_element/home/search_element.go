package home

import (
	"fmt"

	pw "github.com/playwright-community/playwright-go"
)

func (home *HomeElement) GetSearchElementTypeButton(buttonName string) pw.Locator {
	return home.page.Locator(fmt.Sprintf("//*[@id=\"__next\"]//*/span[contains(text(), '%s')]/parent::node()", buttonName))
}

func (home *HomeElement) GetSearchElementCategoryDropdown() pw.Locator {
	return home.page.Locator("//div[contains(@class, 'ui-molecules-select searchBarHome--main__filter_select')]")
}

func (home *HomeElement) GetSearchElementCategoryOption() ([]pw.Locator, error) {
	return home.page.Locator("//ul[@class='ui-molecules-select__field-options ui-molecules-select__field-options-position-bottom']/li").All()
}

func (home *HomeElement) GetSearchElementRegionInput() pw.Locator {
	return home.page.Locator("//input[@name='search-bar']")
}

func (home *HomeElement) GetSearchElementSubmitButton() pw.Locator {
	return home.page.Locator("//button[@class='searchBarHome--submit ui-atomic-button ui-atomic-button__size-default ui-atomic-button__theme-primary']")
}

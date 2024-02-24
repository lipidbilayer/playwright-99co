package home

import (
	"fmt"

	pw "github.com/playwright-community/playwright-go"
)

func (home *HomeElement) GetUpperFooterElementButtonWithButtonName(buttonName string) pw.Locator {
	return home.page.Locator(fmt.Sprintf("//*[@id=\"__next\"]/div[3]/div[1]//*/span[contains(text(), '%s')]", buttonName))
}

func (home *HomeElement) GetUpperFooterElementDivWithButtonName(buttonName string) pw.Locator {
	return home.page.Locator(fmt.Sprintf("//*[@id=\"__next\"]/div[3]/div[1]//*/span[contains(text(), '%s')]/parent::node()/following-sibling::div", buttonName))
}

func (home *HomeElement) GetUpperFooterElementChildUrlWithButtonName(buttonName string) ([]pw.Locator, error) {
	return home.page.Locator(fmt.Sprintf("//*[@id=\"__next\"]/div[3]/div[1]//*/span[contains(text(), '%s')]/parent::node()/following-sibling::div/child::node()", buttonName)).All()
}

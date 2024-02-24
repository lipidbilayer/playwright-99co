package home

import (
	"fmt"

	"99.co/web-test/lib/apperror"
	"github.com/playwright-community/playwright-go"
	pw "github.com/playwright-community/playwright-go"
)

func (home *HomeAction) ISeeFooterButtonPropertyForSale(buttonName string) error {
	button := home.element.GetUpperFooterElementButtonWithButtonName(buttonName)
	return home.assertion.Locator(button).ToBeVisible(pw.LocatorAssertionsToBeVisibleOptions{Timeout: &timeout})
}

func (home *HomeAction) IClickFooterButtonPropertyForSale(buttonName string) error {
	button := home.element.GetUpperFooterElementButtonWithButtonName(buttonName)
	err := button.Click()
	return err
}

func (home *HomeAction) ISeeFooterPropertyForSaleExpanded(buttonName string) error {
	element := home.element.GetUpperFooterElementDivWithButtonName(buttonName)
	return home.assertion.Locator(element).ToBeVisible(pw.LocatorAssertionsToBeVisibleOptions{Timeout: &timeout})
}

func (home *HomeAction) ISeeFooterPropertyForSaleLinks(buttonName string) (int, error) {
	element, err := home.element.GetUpperFooterElementChildUrlWithButtonName(buttonName)
	return len(element), err
}

func (home *HomeAction) IClickFooterPropertyForSaleLinksAndValid(buttonName string, index int) error {
	element, err := home.element.GetUpperFooterElementChildUrlWithButtonName(buttonName)
	if err != nil {
		return err
	}

	url, err := element[index].GetAttribute("href")
	if err != nil {
		return err
	}

	err = element[index].Click(pw.LocatorClickOptions{Button: pw.MouseButtonMiddle})
	if err != nil {
		return err
	}

	newTab, err := home.browserPage.Context().WaitForEvent("page")
	if err != nil {
		return err
	}

	newPage := newTab.(playwright.Page)
	err = newPage.BringToFront()
	if err != nil {
		return err
	}

	err = home.assertion.Page(newPage).ToHaveURL(url)
	if err != nil {
		return err
	}
	res, err := newPage.Reload()
	status := res.Status()
	if status != 200 {
		err = apperror.NewErrorBase(fmt.Sprintf("response return %s", res.StatusText()), apperror.PageInvalid)
	}
	defer newPage.Close()
	return err
}

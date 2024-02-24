package home

import (
	"fmt"
	"regexp"

	"99.co/web-test/lib/apperror"
	"99.co/web-test/lib/helper"
	pw "github.com/playwright-community/playwright-go"
)

func (home *HomeAction) ISeeSearchTypeButton(buttonName string) error {
	button := home.element.GetSearchElementTypeButton(buttonName)
	return home.assertion.Locator(button).ToBeVisible(pw.LocatorAssertionsToBeVisibleOptions{Timeout: &timeout})
}

func (home *HomeAction) IVerifySearchTypeButtonActive(buttonName string) error {
	button := home.element.GetSearchElementTypeButton(buttonName)
	return home.assertion.Locator(button).ToHaveClass(regexp.MustCompile(`\b(?:is[-]active)\b`))
}

func (home *HomeAction) IClickSearchTypeButton(buttonName string) error {
	button := home.element.GetSearchElementTypeButton(buttonName)
	return button.Click()
}

func (home *HomeAction) ISeeSearchCategoryDropdown() error {
	dropdown := home.element.GetSearchElementCategoryDropdown()
	return home.assertion.Locator(dropdown).ToBeVisible(pw.LocatorAssertionsToBeVisibleOptions{Timeout: &timeout})
}

func (home *HomeAction) IClickSearchCategoryDropdown() error {
	dropdown := home.element.GetSearchElementCategoryDropdown()
	return dropdown.Click()
}

func (home *HomeAction) ISeeSearchCategoryDropdownOption() error {
	dropdowns, err := home.element.GetSearchElementCategoryOption()
	if err != nil {
		return err
	}
	for i := 0; i < len(dropdowns); i++ {
		err := home.assertion.Locator(dropdowns[i]).ToBeVisible(pw.LocatorAssertionsToBeVisibleOptions{Timeout: &timeout})
		if err != nil {
			return err
		}
	}
	return nil
}

func (home *HomeAction) IVerifySearchCategoryDropdownOption(option string, index int) error {
	dropdowns, err := home.element.GetSearchElementCategoryOption()
	if err != nil {
		return err
	}
	if index > len(dropdowns) {
		return apperror.NewErrorBase(fmt.Sprintf("dropdown options in valid for index: %d, option: %s", index, option), apperror.ElementInvalid)
	}

	err = home.assertion.Locator(dropdowns[index]).ToBeVisible(pw.LocatorAssertionsToBeVisibleOptions{Timeout: &timeout})
	if err != nil {
		return err
	}
	return home.assertion.Locator(dropdowns[index]).ToHaveText(option)
}

func (home *HomeAction) IClickSearchCategoryDropdownOption(option string) error {
	dropdowns, err := home.element.GetSearchElementCategoryOption()
	if err != nil {
		return err
	}
	locator, err := helper.GetLocatorDropdownOptionWithName(dropdowns, option)
	if err != nil {
		return err
	}
	err = locator.Click()
	if err != nil {
		return err
	}
	return nil
}

func (home *HomeAction) ITypeSearchRegionInput(region string) error {
	input := home.element.GetSearchElementRegionInput()
	err := input.Fill(region)
	if err != nil {
		return err
	}
	return nil
}

func (home *HomeAction) IClickSearchSubmitButton(typeName string, category string, region string) error {
	err := home.element.GetSearchElementSubmitButton().Click()
	if err != nil {
		return err
	}
	expectedUrl := fmt.Sprintf("**/id/%s/%s/%s", typeName, category, region)
	return home.browserPage.WaitForURL(expectedUrl)
}

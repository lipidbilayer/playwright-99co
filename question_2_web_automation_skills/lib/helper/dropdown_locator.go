package helper

import (
	"99.co/web-test/lib/apperror"
	pw "github.com/playwright-community/playwright-go"
)

func GetLocatorDropdownOptionWithName(dropdowns []pw.Locator, option string) (pw.Locator, error) {
	for i := 0; i < len(dropdowns); i++ {
		dropdownText, err := dropdowns[i].TextContent()
		if err != nil {
			return nil, err
		}
		if dropdownText == option {
			return dropdowns[i], nil
		}
	}
	return nil, apperror.NewErrorBase("option not found", apperror.ElementInvalid)
}

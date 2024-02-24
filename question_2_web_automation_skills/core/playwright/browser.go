package playwright

import (
	pw "github.com/playwright-community/playwright-go"
)

func (pe *PlaywrightEngine) OpenBrowser() error {
	browser, err := pe.Chromium.Launch(pw.BrowserTypeLaunchOptions{
		Headless: pw.Bool(false),
	})
	pe.browser = browser

	if err != nil {
		return err
	}

	context, err := browser.NewContext()
	if err != nil {
		return err
	}

	page, err := context.NewPage()
	if err != nil {
		return err
	}
	pe.Page = page
	return nil
}

func (pe *PlaywrightEngine) CloseBrowser() error {
	return pe.browser.Close()
}

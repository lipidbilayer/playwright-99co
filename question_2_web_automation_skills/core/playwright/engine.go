package playwright

import (
	pw "github.com/playwright-community/playwright-go"
)

type PlaywrightEngine struct {
	*pw.Playwright
	browser pw.Browser
	Page    pw.Page
}

func New() (*PlaywrightEngine, error) {
	pwEngine, err := pw.Run()
	if err != nil {
		return nil, err
	}
	playwright := PlaywrightEngine{Playwright: pwEngine}
	return &playwright, nil
}

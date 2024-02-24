package ninenineco_test

import (
	"log"
	"os"
	"testing"

	pw "99.co/web-test/core/playwright"
	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

var (
	pwe     *pw.PlaywrightEngine
	context playwright.BrowserContext
)

var DEFAULT_CONTEXT_OPTIONS = playwright.BrowserNewContextOptions{
	AcceptDownloads: playwright.Bool(true),
	HasTouch:        playwright.Bool(true),
}

func TestMain(m *testing.M) {
	BeforeAll()
	code := m.Run()
	AfterAll()
	os.Exit(code)
}

func BeforeAll() {
	var err error
	pwe, err = pw.New()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
}

func AfterAll() {
	if err := pwe.Stop(); err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
}

func BeforeEach(t *testing.T, contextOptions ...playwright.BrowserNewContextOptions) {
	err := pwe.OpenBrowser()
	require.NoError(t, err)
}

func AfterEach(t *testing.T, closeContext ...bool) {
	// if len(closeContext) == 0 {
	// 	if err := context.Close(); err != nil {
	// 		t.Errorf("could not close context: %v", err)
	// 	}
	// }
}

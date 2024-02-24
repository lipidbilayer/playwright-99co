package searchresult

import (
	"fmt"
	"strings"

	"99.co/web-test/lib/apperror"
)

func (action *SearchResultAction) IVerifySearchResultHeaderTitle(typeName string, category string, region string) error {
	title, err := action.element.GetSearchResultElementHeaderTitle().TextContent()
	if err != nil {
		return err
	}
	//Ruko Dijual di Jawa Barat
	title = strings.ToLower(title)
	expectedTitle := fmt.Sprintf("%s di%s di %s", category, typeName, strings.ReplaceAll(region, "-", " "))
	if title != expectedTitle {
		return apperror.NewErrorBase("search result not match with search parameter", apperror.ResultInvalid)
	}
	return nil
}

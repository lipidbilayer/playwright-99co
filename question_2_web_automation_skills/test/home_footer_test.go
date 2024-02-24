package ninenineco_test

import (
	"testing"

	"99.co/web-test/app/action/home"
	"github.com/stretchr/testify/require"
)

func TestHyperlinkUnderPropertyForSaleFooter(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)

	buttonName := "properti dijual"


	ha := home.New(pwe.Page)

	err := ha.ISeeFooterButtonPropertyForSale(buttonName)
	require.NoError(t, err)
	require.NoError(t, err, "property for sale button should be visible")

	err = ha.IClickFooterButtonPropertyForSale(buttonName)
	require.NoError(t, err, "property for sale button should be clickable")

	err = ha.ISeeFooterPropertyForSaleExpanded(buttonName)
	require.NoError(t, err, "property for sale section should be visible")

	count, err := ha.ISeeFooterPropertyForSaleLinks(buttonName)
	require.NoError(t, err, "property for sale links should be visible")
	require.Equal(t, 15, count)

	for i := 0; i < count; i++ {
		err = ha.IClickFooterPropertyForSaleLinksAndValid(buttonName, i)
		require.NoError(t, err, "links should be open in new tab and return status 200")
	}
}

func TestHyperlinkUnderPropertyForRenFooter(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)

	buttonName := "properti disewa"

	err := pwe.OpenBrowser()
	require.NoError(t, err)
	ha := home.New(pwe.Page)

	err = ha.ISeeFooterButtonPropertyForSale(buttonName)
	require.NoError(t, err)
	require.NoError(t, err, "property for sale button should be visible")

	err = ha.IClickFooterButtonPropertyForSale(buttonName)
	require.NoError(t, err, "property for sale button should be clickable")

	err = ha.ISeeFooterPropertyForSaleExpanded(buttonName)
	require.NoError(t, err, "property for sale section should be visible")

	count, err := ha.ISeeFooterPropertyForSaleLinks(buttonName)
	require.NoError(t, err, "property for sale links should be visible")
	require.Equal(t, 15, count)

	for i := 0; i < count; i++ {
		err = ha.IClickFooterPropertyForSaleLinksAndValid(buttonName, i)
		require.NoError(t, err, "links should be open in new tab and return status 200")
	}
}

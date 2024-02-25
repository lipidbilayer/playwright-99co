package ninenineco_test

import (
	"testing"

	"99.co/web-test/app/action/home"
	searchresult "99.co/web-test/app/action/searchresult"
	"github.com/stretchr/testify/assert"
)

func TestSearchTypeButton(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)

	typesName := []string{"Saya ingin Sewa", "Saya ingin Beli"}
	ha := home.New(pwe.Page)
	for i := 0; i < len(typesName); i++ {
		err := ha.IClickSearchTypeButton(typesName[i])
		assert.NoError(t, err, "Search type button should be clickable", typesName[i])

		err = ha.IVerifySearchTypeButtonActive(typesName[i])
		assert.NoError(t, err, "Search type button should be active when clicked", typesName[i])
	}
}

func TestSearchDropdownOption(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)

	ha := home.New(pwe.Page)

	err := ha.ISeeSearchCategoryDropdown()
	assert.NoError(t, err, "Category Dropdown should be visible")

	err = ha.IClickSearchCategoryDropdown()
	assert.NoError(t, err, "Category Dropdown should be clickable")

	err = ha.ISeeSearchCategoryDropdownOption()
	assert.NoError(t, err, "Category Dropdown option should be visible")

	expectedOptions := []string{"rumah", "apartment", "ruko", "villa", "komersial", "tanah", "kost", "ruang kantor", "gudang", "hotel", "kios", "pabrik", "gedung", "kondotel", "toko"}
	for i := 0; i < len(expectedOptions); i++ {
		err = ha.IVerifySearchCategoryDropdownOption(expectedOptions[i], i)
		assert.NoError(t, err, "Category Dropdown option should be match", expectedOptions[i])
	}
}

func TestSearchBeliRukoJawaBarat(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)

	typeName := "Saya ingin Beli"

	ha := home.New(pwe.Page)

	err := ha.IClickSearchTypeButton(typeName)
	assert.NoError(t, err, "Saya ingin Beli button should be clickable")

	err = ha.IClickSearchCategoryDropdown()
	assert.NoError(t, err, "Category Dropdown should be clickable")

	err = ha.IClickSearchCategoryDropdownOption("ruko")
	assert.NoError(t, err, "Category Dropdown option should be selected")

	err = ha.ITypeSearchRegionInput("Jawa Barat")
	assert.NoError(t, err, "Region input should be typeable")

	err = ha.IClickSearchSubmitButton("jual", "ruko", "jawa-barat")
	assert.NoError(t, err, "Submit button should be clickable")

	sr := searchresult.New(pwe.Page)

	err = sr.IVerifySearchResultHeaderTitle("jual", "rumah", "jawa-barat")
	assert.NoError(t, err, "Search result title should be match with search params")
}

func TestSearchSewaRumahJawaBarat(t *testing.T) {
}

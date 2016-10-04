package stepImpl

import (
	"github.com/apoorvam/gauge-example-go/pages"
	"github.com/getgauge-contrib/gauge-go/gauge"
	m "github.com/getgauge-contrib/gauge-go/models"
)

var _ = gauge.Step("Create a product <table>", func(tbl *m.Table) {
	createProductPage := pages.CreateProductPage{Page: page}
	for _, row := range tbl.Rows {
		title := row.Cells[0]
		description := row.Cells[1]
		author := row.Cells[2]
		price := row.Cells[3]
		gotoNewProductsPage()
		createProductPage.Create(title, description, author, price)
	}
})

func gotoNewProductsPage() {
	driver.Get(pages.NewProductURL)
}

var _ = gauge.Step("On new products page", gotoNewProductsPage)

var _ = gauge.Step("On product page", func() {
	driver.Get(pages.ProductListPageURL)
})

var _ = gauge.Step("Search for product <productName>", func(productName string) {
	productListPage := pages.ProductListPage{Page: page}
	productListPage.Search(productName)
})

var _ = gauge.Step("Open description for product <productName>", func(productName string) {
	productListPage := pages.ProductListPage{Page: page}
	productListPage.OpenFirstProduct()
})

var _ = gauge.Step("Verify product <productAttr> as <authorName>", func(productAttr, authorName string) {
	productPage := pages.ProductPage{Page: page}
	productPage.VerifyProductSpecifier(productPage.GetWebElementByName(productAttr), authorName)
})

var _ = gauge.Step("Check product specifier has new value <table>", func(tbl *m.Table) {
	for _, row := range tbl.Rows {
		productPage := pages.ProductPage{Page: page}
		specifier := productPage.GetWebElementByName(row.Cells[0])
		productPage.VerifyProductSpecifier(specifier, row.Cells[1])
	}
})

var _ = gauge.Step("Delete this product", func() {
	productPage := pages.ProductPage{Page: page}
	productPage.Delete()
})

var _ = gauge.Step("Open product edit page for stored productId", func() {
	editPage := pages.EditProductPage{Page: page}
	editPage.Page.Driver.Get(pages.EditProductURL(fetchStringFromScenarioDataStore("productId")))
})

var _ = gauge.Step("Update product specifier to new value <productData>", func(productData *m.Table) {
	for _, row := range productData.Rows {
		editPage := pages.EditProductPage{Page: page}
		editPage.Page.Driver.Get(pages.EditProductURL(fetchStringFromScenarioDataStore("productId")))
		editPage.UpdateProductValue(row.Cells[0], row.Cells[1])
	}
})

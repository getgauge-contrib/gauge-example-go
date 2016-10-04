package stepImpl

import (
	"github.com/apoorvam/gauge-example-go/pages"
	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/getgauge-contrib/gauge-go/testsuit"
)

func storeStringToScenarioDataStore(key, value string) {
	gauge.GetScenarioStore()[key] = value
}

func storeStringToSpecDataStore(key, value string) {
	gauge.GetSpecStore()[key] = value
}

func storeStringToSuiteDataStore(key, value string) {
	gauge.GetSuiteStore()[key] = value
}

func fetchStringFromScenarioDataStore(key string) string {
	return gauge.GetScenarioStore()[key].(string)
}

func fetchStringFromSpecDataStore(key string) string {
	return gauge.GetSpecStore()[key].(string)
}

func fetchStringFromSuiteDataStore(key string) string {
	return gauge.GetSuiteStore()[key].(string)
}

var _ = gauge.BeforeScenario(func() {
	productListPage := pages.ProductListPage{Page: page}
	productListPage.Page.Driver.Get(pages.ProductListPageURL)
	productListPage.Search("The Way to Go")
	productListPage.OpenFirstProduct()
	productText, _ := productListPage.Page.FindElementByCss(pages.ProductID).Text()
	storeStringToScenarioDataStore("productId", productText)
}, []string{"edit"}, testsuit.AND)

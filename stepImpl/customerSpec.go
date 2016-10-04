package stepImpl

import (
	"github.com/apoorvam/gauge-example-go/pages"
	"github.com/getgauge-contrib/gauge-go/gauge"
	m "github.com/getgauge-contrib/gauge-go/models"
)

var _ = gauge.Step("On the customer page", func() {
	driver.Get(pages.CustomerPageURL)
})

var _ = gauge.Step("Search for customer <customer>", func(customerName string) {
	customerPage := pages.CustomerPage{Page: page}
	customerPage.SearchUser(customerName)
})

var verifyCustomer = gauge.Step("The customer <customerName> is listed", func(customerName string) {
	customerPage := pages.CustomerPage{Page: page}
	customerPage.VerifyUserListed(customerName)
})

var _ = gauge.Step("Search for customers <customersTable>", func(customersTable *m.Table) {
	for _, row := range customersTable.Rows {
		customerPage := pages.CustomerPage{Page: page}
		customerName := row.Cells[0]
		customerPage.SearchUser(customerName)
		customerPage.VerifyUserListed(customerName)
	}
})

var _ = gauge.Step("Just registered customer is listed", func() {
	expectedUser := fetchStringFromScenarioDataStore("currentUser")
	customerPage := pages.CustomerPage{Page: page}
	customerPage.VerifyUserListed(expectedUser)
})

package pages

import (
	"testing"
)

type CustomerPage struct {
	Page Page
}

var (
	CustomerPageURL  = URL + "admin/customers/"
	userNameResult   = "table#index_table_customers tbody tr:nth-child(1) td.col-username"
	userNameID       = "q_username"
	filterButtonName = "commit"
)

func (c *CustomerPage) SearchUser(userName string) {
	c.Page.FindElementById(userNameID).Clear()
	c.Page.FindElementById(userNameID).SendKeys(userName)
	c.Page.FindElementByName(filterButtonName).Click()
}

func (c *CustomerPage) VerifyUserListed(userName string) {
	var t *testing.T
	actualUserName, err := c.Page.FindElementByCss(userNameResult).Text()
	if err != nil {
		t.Errorf("Failed to get element text: %s", err.Error())
	}
	if userName != actualUserName {
		t.Errorf("User %s is not listed", userName)
	}
}

package pages

import (
	"fmt"

	. "github.com/getgauge-contrib/gauge-go/testsuit"
	"github.com/tebeka/selenium"
)

type ProductPage struct {
	Page Page
}

var (
	ProductID           = "#main_content table tbody tr:nth-child(1) td"
	ProductTitle        = "#main_content table tbody tr:nth-child(2) td"
	ProductDescription  = "#main_content table tbody tr:nth-child(3) td"
	ProductAuthor       = "#main_content table tbody tr:nth-child(4) td"
	ProductPrice        = "#main_content table tbody tr:nth-child(5) td"
	ProductDeleteButton = "#titlebar_right div.action_items span.action_item:nth-child(2) a"
)

func (p *ProductPage) GetWebElementByName(elementName string) selenium.WebElement {
	switch elementName {
	case "title":
		return p.Page.FindElementByCss(ProductTitle)
	case "description":
		return p.Page.FindElementByCss(ProductDescription)
	case "author":
		return p.Page.FindElementByCss(ProductAuthor)
	case "price":
		return p.Page.FindElementByCss(ProductPrice)
	}
	return nil
}

func (p *ProductPage) VerifyProductSpecifier(element selenium.WebElement, value string) {
	text, err := element.Text()
	if err != nil {
		T.Fail(fmt.Errorf("Failed to find element text"))
	}
	if text != value {
		T.Fail(fmt.Errorf("want: %s, got: %s", value, text))
	}
}

func (p *ProductPage) Delete() {
	p.Page.FindElementByCss(ProductDeleteButton).Click()
	p.Page.driver().AcceptAlert()
}

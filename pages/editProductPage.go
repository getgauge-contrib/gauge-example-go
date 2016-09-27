package pages

import "strings"

type EditProductPage struct {
	Page Page
}

func EditProductURL(productID string) string {
	return URL + "admin/products/" + productID + "/edit"
}

func (e *EditProductPage) UpdateProductValue(specifier, newValue string) {
	switch strings.ToLower(specifier) {
	case "title":
		e.Page.FindElementById(newProductTitle).Clear()
		e.Page.FindElementById(newProductTitle).SendKeys(newValue)
		break
	case "description":
		e.Page.FindElementById(newProductDescription).Clear()
		e.Page.FindElementById(newProductDescription).SendKeys(newValue)
		break
	case "author":
		e.Page.FindElementById(newProductAuthor).Clear()
		e.Page.FindElementById(newProductAuthor).SendKeys(newValue)
		break
	case "price":
		e.Page.FindElementById(newProductPrice).Clear()
		e.Page.FindElementById(newProductPrice).SendKeys(newValue)
		break
	case "imageFileName":
		e.Page.FindElementById(newProductImage).Clear()
		e.Page.FindElementById(newProductImage).SendKeys(newValue)
		break
	}
	e.Page.FindElementByName(newProductCreateButton).Click()
}

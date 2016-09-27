package pages

type CreateProductPage struct {
	Page Page
}

var (
	NewProductURL = URL + "admin/products/new/"

	newProductTitle = "product_title"

	newProductDescription = "product_description"

	newProductAuthor = "product_author"

	newProductPrice = "product_price"

	newProductCreateButton = "commit"

	newProductImage = "product_image_file_name"
)

func (c *CreateProductPage) Create(title, description, author, price string) {
	c.Page.FindElementById(newProductTitle).SendKeys(title)
	c.Page.FindElementById(newProductDescription).SendKeys(description)
	c.Page.FindElementById(newProductAuthor).SendKeys(author)
	c.Page.FindElementById(newProductPrice).SendKeys(price)
	c.Page.FindElementById(newProductImage).SendKeys("not available")
	c.Page.FindElementByName(newProductCreateButton).Click()
}

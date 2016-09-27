package pages

type ProductListPage struct {
	Page Page
}

var (
	ProductListPageURL = URL + "admin/products/"

	ProductSearchTitle = "q_title"

	ProductSearchCommit = "commit"

	ProductSearchFirstElement = "#main_content table tbody tr:nth-child(1) td.product a"
)

func (p *ProductListPage) Search(name string) {
	p.Page.FindElementById(ProductSearchTitle).SendKeys(name)
	p.Page.FindElementByName(ProductSearchCommit).Click()
}

func (p *ProductListPage) OpenFirstProduct() {
	p.Page.FindElementByCss(ProductSearchFirstElement).Click()
}

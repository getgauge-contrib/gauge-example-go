package stepImpl

import (
	"fmt"
	"os"

	"github.com/apoorvam/gauge-example-go/pages"
	"github.com/getgauge-contrib/gauge-go/gauge"
	. "github.com/getgauge-contrib/gauge-go/testsuit"
	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver
var page pages.Page

var _ = gauge.BeforeSuite(func() {
	var err error
	// set browser as chrome
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	driver, err = selenium.NewRemote(caps, os.Getenv("SELENIUM_HUB_IP"))
	// remote to selenium server
	if err != nil {
		T.Fail(fmt.Errorf("Failed to open session: %s\n", err.Error()))
	}
	page = pages.Page{Driver: driver}
	err = driver.Get(pages.URL)
	if err != nil {
		T.Fail(fmt.Errorf("Failed to load page: %s\n", err.Error()))
		return
	}
}, []string{}, AND)

var _ = gauge.AfterSuite(func() {
	driver.Quit()
}, []string{}, AND)

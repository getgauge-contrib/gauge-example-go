package stepImpl

import (
	"os"
	"testing"

	"github.com/apoorvam/gauge-example-go/pages"
	"github.com/manuviswam/gauge-go/gauge"
	"github.com/manuviswam/gauge-go/testsuit"
	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver
var page pages.Page
var t *testing.T

var _ = gauge.BeforeSuite(func() {
	var err error
	// set browser as chrome
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	driver, err = selenium.NewRemote(caps, os.Getenv("SELENIUM_HUB_IP"))
	// remote to selenium server
	if err != nil {
		t.Errorf("Failed to open session: %s\n", err.Error())
	}
	page = pages.Page{Driver: driver}
	err = driver.Get(pages.URL)
	if err != nil {
		t.Errorf("Failed to load page: %s\n", err.Error())
		return
	}
}, []string{}, testsuit.AND)

var _ = gauge.AfterSuite(func() {
	driver.Quit()
}, []string{}, testsuit.AND)

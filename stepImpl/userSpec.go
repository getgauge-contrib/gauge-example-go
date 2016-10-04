package stepImpl

import (
	"math/rand"
	"time"

	"github.com/apoorvam/gauge-example-go/pages"
	"github.com/getgauge-contrib/gauge-go/gauge"
)

var _ = gauge.Step("On signup page", func() {
	driver.Get(pages.SignupPageURL)
})

var _ = gauge.Step("Fill in and send registration form", func() {
	signupPage := pages.SignupPage{Page: page}
	randomName := getRandomName(8)
	signupPage.Page.FindElementById(pages.UsernameID).SendKeys(randomName)
	signupPage.Page.FindElementById(pages.UsernameEmailID).SendKeys(randomName + "@domain.com")
	signupPage.Page.FindElementById(pages.UserPasswordID).SendKeys("qweqwe")
	signupPage.Page.FindElementById(pages.UserPasswordConID).SendKeys("qweqwe")
	signupPage.Page.FindElementByName(pages.UserCommitName).Click()
	storeStringToScenarioDataStore("currentUser", randomName)
})

func getRandomName(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

package main

import (
	"github.com/sclevine/agouti"
	"log"
	"time"
)

func main() {
	username := "user"
	password := "pass"
	website := "web"

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver : %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("failed to open page : %v", err)
	}

	err = page.Navigate(website)
	if err != nil {
		log.Fatalf("failed to navigate the site : %v", err)
	}

	loginLink := page.FindByClass("link-login")
	log.Println(loginLink)
	err = loginLink.Click()
	if err != nil {
		log.Fatalf("failed to click the login link : %v", err)
	}

	time.Sleep(time.Second * 5)

	loginTab := page.Find(".login-tab.login-tab-r")
	log.Println(loginTab)
	err = loginTab.Click()
	if err != nil {
		log.Fatalf("failed to click the login tab ： %v", err)
	}

	time.Sleep(time.Second * 5)

	userNameEntry := page.FindByID("loginname")
	userNameEntry.SendKeys(username)

	userPassEntry := page.FindByID("nloginpwd")
	userPassEntry.SendKeys(password)

	loginSubmit := page.FindByID("loginsubmit")
	loginSubmit.Click()

	time.Sleep(time.Second * 5)

	searchTextInput := page.FindByID("key")
	log.Println(searchTextInput)
	err = searchTextInput.SendKeys("数据线")
	if err != nil {
		log.Fatalf("failed to input the  search text.  ： %v", err)
	}

	searchButton := page.FindByClass("button")
	log.Println(searchButton)
	err = searchButton.Click()
	if err != nil {
		log.Fatalf("failed to click the search button ： %v", err)
	}

	time.Sleep(time.Second * 10)

	page.Refresh()

	time.Sleep(time.Second * 10)

}

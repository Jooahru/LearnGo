package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/will/learngo/scrapper"
)

const fileName string = "jobs.csv"

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName,fileName)
}

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

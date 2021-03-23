package main

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/woojebiz/learngo/scrapper"
)

const fileName string = "jobs.csv"

func main() {
	//scrapper.Scrape("go")

	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	e.GET("/", handleHome)

	// scrape
	e.POST("/scrape", handleScrape)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}

// Handler home
func handleHome(c echo.Context) error {
	return c.File("home.html")
	// return c.String(http.StatusOK, "Hello, World!")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

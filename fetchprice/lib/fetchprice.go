package lib

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// conditions
const (
	used       = "used"
	new        = "new"
	complete   = "complete"
	graded     = "graded"
	boxOnly    = "boxOnly"
	manualOnly = "manualOnly"
)

func FetchPrice(name, platform, condition, edition string) (price float64, curr string) {
	curr = "USD"
	price = -1

	// generate URL
	url := getPCUrl(name, platform, edition)
	query := getPCQuery(condition)

	// setup scraper
	c := colly.NewCollector(
		colly.AllowedDomains("www.pricecharting.com"),
	)

	// scrape data
	c.OnHTML(query, func(h *colly.HTMLElement) {
		priceStr := strings.Trim(h.Text, "$ \t\n")

		var err error
		price, err = strconv.ParseFloat(priceStr, 32)
		if err != nil {
			fmt.Println("error parsing value")
		}
	})

	// display URL
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Sending request: %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err.Error())
	})

	// kick of request
	c.Visit(url)

	return price, curr
}

func getPCQuery(condition string) string {
	// price charting element ids
	const (
		PC_usedPrice     = "used_price"
		PC_completePrice = "complete_price"
		PC_newPrice      = "new_price"
		PC_gradedPrice   = "graded_price"
		PC_boxOnly       = "box_only_price"
		PC_manualOnly    = "manual_only_price"
	)
	var priceElementId string

	switch condition {
	case used:
		priceElementId = PC_usedPrice
	case complete:
		priceElementId = PC_completePrice
	case new:
		priceElementId = PC_newPrice
	case graded:
		priceElementId = PC_gradedPrice
	case boxOnly:
		priceElementId = PC_boxOnly
	case manualOnly:
		priceElementId = PC_manualOnly
	default:
		priceElementId = PC_usedPrice
	}

	return fmt.Sprintf("td[id=%s] .price", priceElementId)
}
func getPCUrl(game, platform, edition string) string {
	g := game
	p := platform
	e := edition

	// TODO: map VGL platform codes to price charting
	// TODO: normalize game name and edition to price charting standard

	var urlResult string
	if strings.TrimSpace(e) == "" {
		urlResult = fmt.Sprintf("https://www.pricecharting.com/game/%s/%s", p, g)
	} else {
		urlResult = fmt.Sprintf("https://www.pricecharting.com/game/%s/%s-%s", p, g, e)
	}

	return urlResult

}

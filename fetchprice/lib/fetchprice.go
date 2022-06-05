package lib

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// conditions
const (
	used       = "used"     // loose, or media only
	new        = "new"      // sealed
	complete   = "complete" // box + instructions + media
	graded     = "graded"
	boxOnly    = "boxOnly"    // box
	manualOnly = "manualOnly" // instructions
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

func NormalizeString(s string) string {
	normalized := strings.TrimSpace(s)
	normalized = strings.ReplaceAll(normalized, ".", "")
	normalized = strings.ReplaceAll(normalized, "'", "%27s")
	normalized = strings.ReplaceAll(normalized, " ", "-")
	normalized = strings.ReplaceAll(normalized, ":", "")
	normalized = strings.ReplaceAll(normalized, "*", "")
	normalized = strings.ReplaceAll(normalized, "`", "")
	normalized = strings.ToLower(normalized)

	return normalized
}

func pcNormalizeString(s string) string {
	normalized := strings.TrimSpace(s)
	normalized = strings.ReplaceAll(normalized, ".", "")
	normalized = strings.ReplaceAll(normalized, "'", "%27s")
	normalized = strings.ReplaceAll(normalized, " ", "-")
	normalized = strings.ReplaceAll(normalized, ":", "")
	normalized = strings.ToLower(normalized)

	return normalized
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

	var pcPlatforms = make(map[string]string)

	pcPlatforms["2600"] = "atari-2600"
	pcPlatforms["3ds"] = "nintendo-3ds"
	pcPlatforms["coleco"] = "colecovision"
	pcPlatforms["dc"] = "sega-dreamcast"
	pcPlatforms["ds"] = "nintendo-ds"
	pcPlatforms["fami"] = "famicom"
	pcPlatforms["famids"] = "famicom-disk-system"
	pcPlatforms["gba"] = "gameboy-advance" // "jp-gameboy-advance"
	pcPlatforms["gbc"] = "gameboy-color"
	pcPlatforms["gb"] = "gameboy"
	pcPlatforms["gc"] = "gamecube"
	pcPlatforms["gen"] = "sega-genesis"
	pcPlatforms["nes"] = "nes"
	pcPlatforms["n64"] = "nintendo-64"
	pcPlatforms["n64dd"] = "jp-nintendo-64"
	pcPlatforms["ps1"] = "playstation"
	pcPlatforms["ps2"] = "playstation-2"
	pcPlatforms["ps3"] = "playstation-3"
	pcPlatforms["ps4"] = "playstation-4"
	pcPlatforms["ps5"] = "playstation-5"
	pcPlatforms["psp"] = "psp"
	pcPlatforms["saturn"] = "saturn"
	pcPlatforms["segacd"] = "sega-cd"
	pcPlatforms["sfc"] = "super-famicom"
	pcPlatforms["snes"] = "super-nintendo"
	pcPlatforms["switch"] = "nintendo-switch"
	pcPlatforms["tgfx16"] = "turbografx-16"
	pcPlatforms["tgfxcd"] = "turbografx-cd"
	pcPlatforms["vita"] = "playstation-vita"
	pcPlatforms["vb"] = "virtual-boy"
	pcPlatforms["wii"] = "wii"
	pcPlatforms["wiiu"] = "wii-u"
	pcPlatforms["wsc"] = "wonderswan-color"
	pcPlatforms["xb"] = "xbox"
	pcPlatforms["xb360"] = "xbox-360"
	pcPlatforms["xb1"] = "xbox-one"
	pcPlatforms["xbx"] = "xbox-series-x"

	p := pcPlatforms[platform]
	g := pcNormalizeString(game)
	e := pcNormalizeString(edition)

	var urlResult string
	if strings.TrimSpace(e) == "" {
		urlResult = fmt.Sprintf("https://www.pricecharting.com/game/%s/%s", p, g)
	} else {
		urlResult = fmt.Sprintf("https://www.pricecharting.com/game/%s/%s-%s", p, g, e)
	}

	return urlResult
}

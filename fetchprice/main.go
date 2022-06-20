package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cocobokostudios/vgl-tools/fetchprice/lib"
	"github.com/gocolly/colly/v2"
)

type Game struct {
	Title        string
	Url          string
	PlatformId   string
	RegionId     string
	Box          string
	Instructions string
	Media        string
	Sealed       string
	Graded       string
	Reproduction string
	PriceData    GamePrice
}

type GamePrice struct {
	Used        float64
	Complete    float64
	New         float64
	Graded      float64
	BoxOnly     float64
	ManualOnly  float64
	RetrievedOn time.Time
}

// conditions
const (
	USED        = "used"     // loose, or media only
	NEW         = "new"      // sealed
	COMPLETE    = "complete" // box + instructions + media
	GRADED      = "graded"
	BOX_ONLY    = "boxOnly"    // box
	MANUAL_ONLY = "manualOnly" // instructions
)

// true, false from CSV
const (
	YES            = "y"
	NO             = "n"
	NOT_APPLICABLE = "na"
)

func (g *Game) GetPrice() float64 {
	var price float64

	if g.Graded == YES {
		// graded
		price = g.PriceData.Graded
	} else if g.Sealed == YES {
		// sealed
		price = g.PriceData.New
	} else if g.Media == YES && g.Box == YES && g.Instructions == YES {
		// complete
		price = g.PriceData.Complete
	} else if g.Box == YES && g.Instructions == YES {
		// box + instructions
		price = g.PriceData.BoxOnly + g.PriceData.ManualOnly
	} else if g.Media == YES && g.Box == YES {
		// box + media
		price = g.PriceData.BoxOnly + g.PriceData.Used
	} else if g.Media == YES {
		// just game
		price = g.PriceData.Used
	} else if g.Box == YES {
		// just box
		price = g.PriceData.BoxOnly
	} else if g.Instructions == YES {
		// just instruction
		price = g.PriceData.ManualOnly
	} else {
		price = 0
	}

	return price
}

func (g *Game) ToStringArr() []string {
	return []string{
		g.Title,
		g.Url,
		g.PlatformId,
		g.RegionId,
		g.Box,
		g.Instructions,
		g.Media,
		g.Sealed,
		g.Graded,
		g.Reproduction,
		string(strings.TrimSpace(fmt.Sprintf("%5.2f", g.GetPrice()))),
	}
}

func parsePrice(text string) float64 {
	priceStr := strings.Trim(text, "$ \t\n")

	var err error
	var price float64

	price, err = strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("error parsing value")
	}

	return price
}

func FetchPrice(url string) GamePrice {
	// create price object
	var gamePrice GamePrice
	gamePrice.RetrievedOn = time.Now()

	// setup scraper
	c := colly.NewCollector(
		colly.AllowedDomains("www.pricecharting.com"),
	)

	// scrape price data
	c.OnHTML(lib.GetPCSelector(USED), func(h *colly.HTMLElement) {
		gamePrice.Used = parsePrice(h.Text)
	})
	c.OnHTML(lib.GetPCSelector(NEW), func(h *colly.HTMLElement) {
		gamePrice.New = parsePrice(h.Text)
	})
	c.OnHTML(lib.GetPCSelector(COMPLETE), func(h *colly.HTMLElement) {
		gamePrice.Complete = parsePrice(h.Text)
	})
	c.OnHTML(lib.GetPCSelector(GRADED), func(h *colly.HTMLElement) {
		gamePrice.Graded = parsePrice(h.Text)
	})
	c.OnHTML(lib.GetPCSelector(BOX_ONLY), func(h *colly.HTMLElement) {
		gamePrice.BoxOnly = parsePrice(h.Text)
	})
	c.OnHTML(lib.GetPCSelector(MANUAL_ONLY), func(h *colly.HTMLElement) {
		gamePrice.ManualOnly = parsePrice(h.Text)
	})

	// display URL
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Sending request: %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err.Error())
	})

	// kick off request
	c.Visit(url)

	return gamePrice
}

func main() {
	// read the CSV
	file, err := os.Open("./data/collection.csv")
	if err != nil {
		fmt.Printf("An error occured while opening the file. Error: %s", err.Error())
	}
	defer file.Close()

	// create new file
	newFile, err := os.Create("./data/collection.prices.csv")
	if err != nil {
		fmt.Printf("An error occured while creating new file. Error: %s", err.Error())
	}
	defer file.Close()

	// read the values
	csvReader := csv.NewReader(file)
	csvReader.TrimLeadingSpace = true
	games, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// write updated file
	csvWriter := csv.NewWriter(newFile)
	var game Game

	for i, line := range games {
		var values [13]string
		if i > 0 {
			fmt.Println(line)
			for j, field := range line {
				if j == 0 {
					values[0] = strings.TrimSpace(field) // title
					game.Title = values[0]
				} else if j == 1 {
					values[1] = field // url
					game.Url = values[1]
				} else if j == 2 {
					values[2] = field // platformId
					game.PlatformId = values[2]
				} else if j == 3 {
					values[3] = field // regionId
					game.RegionId = values[3]
				} else if j == 4 {
					values[4] = field // box
					game.Box = values[4]
				} else if j == 5 {
					values[5] = field // instruction
					game.Instructions = values[5]
				} else if j == 6 {
					values[6] = field // media
					game.Media = values[6]
				} else if j == 7 {
					values[7] = field // sealed
					game.Sealed = values[7]
				} else if j == 8 {
					values[8] = field // graded
					game.Graded = values[8]
				} else if j == 9 {
					values[9] = field // reproduction
					game.Reproduction = values[9]
				}
			}

			game.PriceData = FetchPrice(game.Url)
			//price, curr := lib.FetchPrice(values[1], values[3], "complete", values[2])
			//gameWithPrice := append(values[:], strings.TrimSpace(fmt.Sprintf("%5.2f", price)), curr)
			csvWriter.Write(game.ToStringArr())
		} else { // header line
			line = append(line, "price")
			csvWriter.Write(line)
		}

		csvWriter.Flush()
	}
}

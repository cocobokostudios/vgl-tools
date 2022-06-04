package main

import (
	"fmt"

	"github.com/cocobokostudios/vgl-tools/fetchprice/lib"
)

func main() {
	price, currency := lib.FetchPrice("Uncharted Waters", "snes", "used", "")
	fmt.Printf("$ %5.2f %s\n", price, currency)

	price, currency = lib.FetchPrice("Uncharted Waters: New Horizons", "snes", "used", "")
	fmt.Printf("$ %5.2f %s\n", price, currency)

	price, currency = lib.FetchPrice("SimCity", "snes", "new", "")
	fmt.Printf("$ %5.2f %s\n", price, currency)

	price, currency = lib.FetchPrice("Final Fantasy", "wsc", "new", "")
	fmt.Printf("$ %5.2f %s\n", price, currency)

	price, currency = lib.FetchPrice("Spidernan: Miles Morales", "ps5", "used", "Ultimate Launch Edition")
	fmt.Printf("$ %5.2f %s\n", price, currency)

}

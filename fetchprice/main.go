package main

import (
	"fmt"

	"github.com/cocobokostudios/vgl-tools/fetchprice/lib"
)

func main() {
	price, currency := lib.FetchPrice("uncharted-waters", "super-nintendo", "used", "")
	fmt.Printf("$ %5.2f %s\n", price, currency)

	price, currency = lib.FetchPrice("simcity", "super-nintendo", "new", "")
	fmt.Printf("$ %5.2f %s\n", price, currency)

	price, currency = lib.FetchPrice("simcity-4", "pc-games", "new", "deluxe-edition")
	fmt.Printf("$ %5.2f %s\n", price, currency)
}

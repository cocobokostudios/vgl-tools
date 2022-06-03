package main

import (
	"fmt"

	"github.com/cocobokostudios/vgl-tools/fetchprice/lib"
)

func main() {
	price, currency := lib.FetchPrice("uncharted-waters", "super-nintendo", "used", "")
	fmt.Printf("$ %5.2f %s\n", price, currency)
}

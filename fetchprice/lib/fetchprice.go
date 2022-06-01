package lib

import (
	"strings"
)

func GetPrice(id string) string {
	return "123"
}

func GetMessage() string {
	message := []string{"hello", "from", "the", "fetchprice", "library"}
	return strings.Join(message, " ")
}

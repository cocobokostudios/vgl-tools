package lib

import (
	"strings"
)

func GetPrice(id string) string {
	return "123"
}

func GetMessage() string {
	message := []string{"hello", "from", "the", "price", "library"}
	return strings.Join(message, " ")
}

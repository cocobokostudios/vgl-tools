package price_lib

import (
	"strings"
)

func GetMessage() string {
	message := []string{"hello", "from", "price"}
	return strings.Join(message, " ")
}

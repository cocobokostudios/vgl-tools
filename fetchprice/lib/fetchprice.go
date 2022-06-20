package lib

import (
	"fmt"
	"strings"
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

func GenerateGameId(title, edition string) string {
	var gameId string

	if edition != "" {
		gameId = fmt.Sprintf("%s-%s", NormalizeString(title), NormalizeString(edition))
	} else {
		gameId = NormalizeString(title)
	}

	return gameId
}

func NormalizeString(s string) string {
	normalized := strings.TrimSpace(s)
	normalized = strings.ReplaceAll(normalized, ".", "")
	normalized = strings.ReplaceAll(normalized, "'", "")
	normalized = strings.ReplaceAll(normalized, " ", "-")
	normalized = strings.ReplaceAll(normalized, ":", "")
	normalized = strings.ReplaceAll(normalized, "*", "")
	normalized = strings.ReplaceAll(normalized, "`", "")
	normalized = strings.ToLower(normalized)

	return normalized
}

func GetPCSelector(condition string) string {
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

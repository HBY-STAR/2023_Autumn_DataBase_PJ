package json

import (
	"testing"
)

func TestJson(t *testing.T) {
	_ = parseUsers()
	_ = parseSeller()
	_ = parseCommodity()
	_ = parsePlatform()
	_ = parseAdmin()

	_ = parseFavorite()
	_ = parseItem()
	_ = parsePriceChange()
	_ = parseMessage()

}

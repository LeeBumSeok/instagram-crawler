package instagram_test

import (
	"GO/dataservice/instagram"
	"testing"
)

func TestPageParserGenerator(t *testing.T) {
	parser := instagram.PageParserGenerator("fff")
	for i := 0; i < 10; i++ {
		page, err := parser()
		if err != nil {
			t.Error(err)
		}
		shortcodes := page.Shortcodes()
		if 0 < len(shortcodes) {
			t.Error(shortcodes)
		}
	}
}

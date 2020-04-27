package collector

import (
	"fmt"
	"testing"
)

func TestCollector(t *testing.T) {
	crawler, _ := NewCollector()
	err := crawler.Get("/trademarkInfoSearchService/applicationNumberSearchInfo")
	fmt.Println(err)
}

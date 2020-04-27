package crawlers

import (
	"fmt"
	"testing"
)

func TestCrawler(t *testing.T) {
	crawler, _ := NewCrawler("http://plus.kipris.or.kr/openapi/rest")
	err := crawler.Get("/trademarkInfoSearchService/applicationNumberSearchInfo")
	fmt.Println(err)
}

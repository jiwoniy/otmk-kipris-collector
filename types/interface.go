package types

import (
	"github.com/jinzhu/gorm"
	"github.com/jiwoniy/otmk-kipris-collector/model"
)

type Collector interface {
	GetEndpoint() string
	GetAccessKey() string
	GetParser() Parser
	GetStorage() Storage
	Get(url string, params map[string]string) ([]byte, error)
	StartCrawler(year string)
	CrawlerApplicationNumber(applicationNumber string) bool

	CreateApplicationNumberList(year string, length int, startNumber int) []string
	CreateApplicationNumber(productCode string, year string, serialNumber int) string

	// for find application number. but it is useless
	GetLastApplicationNumber(startNumber string, LastNumber string, checker func(string) bool) (string, string, error)
	GetMidValue(startNumber int, lastNumber int) int
	IsTestApplicationNumberExist(answer string) func(string) bool
}

type Query interface {
	GetApplicationNumber(applicationNumber string) *model.TradeMarkInfo
}

type Parser interface {
	Read(filename string) ([]byte, error)
	Print(v interface{})
	Parse(data []byte, v interface{}) error
}

type Storage interface {
	GetDB() *gorm.DB
	CloseDB()
	Create(v Model) error
	GetKiprisApplicationNumber(v model.KiprisApplicationNumber, data *model.KiprisApplicationNumber)
	GetKiprisApplicationNumberList(v model.KiprisApplicationNumber, data *[]model.KiprisApplicationNumber)
	GetKiprisCollector(v model.KiprisCollectorStatus, data *model.KiprisCollectorStatus)
	GetTradeMarkInfo(v model.TradeMarkInfo, data *model.TradeMarkInfo)
	GetTrademarkDesignationGoodstInfo(v model.TrademarkDesignationGoodstInfo, data *[]model.TrademarkDesignationGoodstInfo)
}

type Model interface {
	Valid() bool
}

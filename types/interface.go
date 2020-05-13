package types

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jiwoniy/otmk-kipris-collector/model"
	"github.com/jiwoniy/otmk-kipris-collector/pagination"
)

type RestHandler func(ctx *gin.Context)

type RestMethod struct {
	Path    string
	Handler RestHandler
}

type TaskParameters struct {
	ProductCode       string
	Year              string
	SerialNumberRange string
	Page              int
	Size              int
}

type RestClient interface {
	GetMethods() ([]RestMethod, error)
	PostMethods() ([]RestMethod, error)
}

type Collector interface {
	// for rest client
	GetMethods() ([]RestMethod, error)
	PostMethods() ([]RestMethod, error)

	GetEndpoint() string
	GetAccessKey() string
	GetParser() Parser
	GetStorage() Storage
	Get(url string, params map[string]string) ([]byte, error)

	// task
	CreateTask(args TaskParameters) error
	CreatManualTask(args TaskParameters) error
	GetTaskList(page int, size int) (*pagination.Paginator, error)
	// GetTaskApplicationNumberList(taskId uint, page int, size int) (*pagination.Paginator, error)

	// crawler
	StartCrawler(taskId uint) error
	CrawlerApplicationNumber(tx *gorm.DB, applicationNumber string) bool

	// collector helper
	// GetApplicationNumberList(args TaskParameters) (*pagination.Paginator, error)

	// CreateApplicationNumberList(year string, length int, startNumber int) []string
	// CreateApplicationNumber(productCode string, year string, serialNumber int) string
	// GetYearLastApplicationNumber(year string) string

	// for find application number. but it is useless
	// GetLastApplicationNumber(startNumber string, LastNumber string, checker func(string) bool) (string, string, error)

	// GetMidValue(startNumber int, lastNumber int) int
	// IsTestApplicationNumberExist(answer string) func(string) bool
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
	CreateTask(applicationNumbers *[]model.KiprisApplicationNumber) error

	GetTaskList(page int, size int) (*pagination.Paginator, error)
	GetTaskApplicationNumberList(tx *gorm.DB, taskId uint, pagination ...int) (*pagination.Paginator, error)

	// GetKiprisApplicationNumber(v model.KiprisApplicationNumber, data *model.KiprisApplicationNumber)
	// GetKiprisApplicationNumberList(v model.KiprisApplicationNumber, data *[]model.KiprisApplicationNumber, startSerialNumber int, endSerialNumber int, page int, size int) (*pagination.Paginator, error)
	GetKiprisCollector(v model.KiprisCollectorStatus, data *model.KiprisCollectorStatus)
	GetTradeMarkInfo(v model.TradeMarkInfo, data *model.TradeMarkInfo)
	GetTrademarkDesignationGoodstInfo(v model.TrademarkDesignationGoodstInfo, data *[]model.TrademarkDesignationGoodstInfo)

	GetYearLastApplicationNumber(year string) string
}

type Model interface {
	Valid() bool
}

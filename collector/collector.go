package collector

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jiwoniy/otmk-kipris-collector/model"
	"github.com/jiwoniy/otmk-kipris-collector/pagination"
	"github.com/jiwoniy/otmk-kipris-collector/parser"
	"github.com/jiwoniy/otmk-kipris-collector/storage"
	"github.com/jiwoniy/otmk-kipris-collector/types"
	"github.com/jiwoniy/otmk-kipris-collector/utils"
)

type kiprisCollector struct {
	endpt     string
	accessKey string
	parser    types.Parser
	storage   types.Storage
	query     types.Query
}

var collectLogger *log.Logger
var taskNumber = 1000

func init() {
	fpLog, err := os.OpenFile("collector_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	collectLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewCollector(config types.CollectorConfig) (types.Collector, error) {
	parserInstance, err := parser.NewParser("xml")

	if err != nil {
		return nil, err
	}

	storage, err := storage.NewStorage(types.StorageConfig{
		DbType:       config.DbType,
		DbConnString: config.DbConnString,
	})

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &kiprisCollector{
		endpt:     config.Endpoint,
		accessKey: config.AccessKey,
		parser:    parserInstance,
		storage:   storage,
	}, nil
}

func (c *kiprisCollector) GetEndpoint() string {
	return c.endpt
}

func (c *kiprisCollector) GetAccessKey() string {
	return c.accessKey
}

func (c *kiprisCollector) GetParser() types.Parser {
	return c.parser
}

func (c *kiprisCollector) GetStorage() types.Storage {
	return c.storage
}

func (c *kiprisCollector) Get(url string, params map[string]string) ([]byte, error) {
	caller, err := utils.BuildRESTCaller(c.endpt).Build()
	if err != nil {
		return nil, err
	}

	body, err := caller.Get(nil, url, params)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// For rest
func (c *kiprisCollector) GetMethods() ([]types.RestMethod, error) {
	restMethods := make([]types.RestMethod, 1)
	restMethods = append(restMethods,
		types.RestMethod{
			Path: "/search",
			Handler: func(ctx *gin.Context) {
				// param := types.TaskParameters{
				// 	ProductCode:       "40",
				// 	Year:              "2017",
				// 	SerialNumberRange: "1,20",
				// }
				// c.GetApplicationNumberList(param)
			},
		},
		types.RestMethod{
			Path: "/ping",
			Handler: func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "pong")
			},
		},
		types.RestMethod{
			Path: "/applicationNumbers",
			Handler: func(ctx *gin.Context) {
				// params := types.TaskParameters{
				// 	ProductCode:       "40",
				// 	Year:              "2017",
				// 	SerialNumberRange: "1,20",
				// }
				// pagination, err := c.GetApplicationNumberList(params)
				// fmt.Println(pagination.Data)
				// fmt.Println(err)
			},
		},
	)
	return restMethods, nil
}

// Kipris Task
func (c *kiprisCollector) CreateTask(args types.TaskParameters) error {
	// 상표법 개정에 따라 서비스표(41), 상표/서비스표(45)는 2016년 9월 1일 이후 출원건에 대해 상표(40)에 통합 되었습니다.
	productCodeList := []string{
		"40",
		"41",
		"45",
	}

	year := args.Year
	serialNumberRange := args.SerialNumberRange

	// require
	if year == "" {
		return fmt.Errorf("parameters's year require")
	}
	// require
	if serialNumberRange == "" {
		yearLastApplicationNumber := c.storage.GetYearLastApplicationNumber(year)
		if yearLastApplicationNumber == "" {
			serialNumberRange = "1," + strconv.Itoa(taskNumber)
		}
	}

	kiprisApplicationNumbers := make([]model.KiprisApplicationNumber, 0)

	yearInt, _ := strconv.Atoi(year)
	if yearInt > 2016 {
		serialNumberList := getSerialNumberList(serialNumberRange)
		for _, serialNumber := range serialNumberList {
			kiprisApplicationNumber := model.KiprisApplicationNumber{
				ApplicationNumber: fmt.Sprintf("%s%s%07d", "40", year, serialNumber),
				ProductCode:       "40",
				Year:              year,
				SerialNumber:      serialNumber,
			}
			kiprisApplicationNumbers = append(kiprisApplicationNumbers, kiprisApplicationNumber)
		}
	} else {
		serialNumberList := getSerialNumberList(serialNumberRange)
		for _, productCode := range productCodeList {
			for _, serialNumber := range serialNumberList {
				kiprisApplicationNumber := model.KiprisApplicationNumber{
					ApplicationNumber: fmt.Sprintf("%s%s%07d", productCode, year, serialNumber),
					ProductCode:       productCode,
					Year:              year,
					SerialNumber:      serialNumber,
				}
				kiprisApplicationNumbers = append(kiprisApplicationNumbers, kiprisApplicationNumber)
			}
		}
	}

	return c.storage.CreateTask(&kiprisApplicationNumbers)
}

// Kipris Manual Task
func (c *kiprisCollector) CreatManualTask(args types.TaskParameters) error {
	productCode := args.ProductCode
	year := args.Year
	serialNumberRange := args.SerialNumberRange

	// require
	if productCode == "" || year == "" || serialNumberRange == "" {
		return fmt.Errorf("Please pass task parameters")
	}

	kiprisApplicationNumbers := make([]model.KiprisApplicationNumber, 0)
	serialNumberList := getSerialNumberList(serialNumberRange)
	for _, serialNumber := range serialNumberList {
		kiprisApplicationNumber := model.KiprisApplicationNumber{
			ApplicationNumber: fmt.Sprintf("%s%s%07d", productCode, year, serialNumber),
			ProductCode:       "40",
			Year:              year,
			SerialNumber:      serialNumber,
		}
		kiprisApplicationNumbers = append(kiprisApplicationNumbers, kiprisApplicationNumber)
	}

	return c.storage.CreateTask(&kiprisApplicationNumbers)
}

func (c *kiprisCollector) GetTaskList(pageParam int, sizeParam int) (*pagination.Paginator, error) {
	pagination, err := c.storage.GetTaskList(pageParam, sizeParam)
	return pagination, err
}

// func (c *kiprisCollector) GetTaskApplicationNumberList(taskId uint, pageParam int, sizeParam int) (*pagination.Paginator, error) {
// 	pagination, err := c.storage.GetTaskApplicationNumberList(taskId, pageParam, sizeParam)
// 	return pagination, err
// }

// func (c *kiprisCollector) GetApplicationNumberList(args types.TaskParameters) (*pagination.Paginator, error) {
// 	searchResult := make([]model.KiprisApplicationNumber, 0)
// 	searchData := model.KiprisApplicationNumber{
// 		Year:        args.Year,
// 		ProductCode: args.ProductCode,
// 	}

// 	startSerialNumber := 0
// 	endSerialNumber := 0
// 	serialRangeList := strings.Split(args.SerialNumberRange, ",")
// 	if len(serialRangeList) == 2 {
// 		startSerialNumber, _ = strconv.Atoi(serialRangeList[0])
// 		endSerialNumber, _ = strconv.Atoi(serialRangeList[1])
// 	}

// 	pagination, err := c.storage.GetKiprisApplicationNumberList(searchData, &searchResult, startSerialNumber, endSerialNumber, args.Page, args.Size)
// 	return pagination, err
// }

func (c *kiprisCollector) StartCrawler(taskId uint) error {
	db := c.storage.GetDB()
	pageSize := 50

	currentTask := model.KiprisTask{}

	if err := db.Table("kipris_tasks").Where("id = ?", taskId).First(&currentTask).Error; err != nil {
		collectLogger.Printf("[StartCrawler Task Id] %d Step 1 - Get Task (error: %s)", taskId, err)
		return fmt.Errorf("[StartCrawler Task Id] %d Step 1 - Get Task", taskId)
	}

	if err := db.Model(&currentTask).Update("started", time.Now()).Error; err != nil {
		collectLogger.Printf("[StartCrawler Task Id] %d Step 2 - Update Task started Date (error: %s)", taskId, err)
		return fmt.Errorf("[StartCrawler Task Id] %d Step 2 - Update Task started Date", taskId)
	}

	return db.Transaction(func(tx *gorm.DB) error {
		pagination, err := c.storage.GetTaskApplicationNumberList(tx, taskId, 1, pageSize)

		if err != nil {
			collectLogger.Printf("[StartCrawler Task Id] %d Step 3 - Get Task Application Number (error: %s)", taskId, err)
			return fmt.Errorf("[StartCrawler Task Id] %d Step 3 - Get Task Application Number", taskId)
		}

		for currentPage := pagination.Page; currentPage <= pagination.TotalPage; currentPage++ {
			data := pagination.Data.(*[]model.KiprisApplicationNumber)
			var wg sync.WaitGroup
			for _, application := range *data {
				wg.Add(1)
				go func(appNumber string) {
					defer wg.Done()
					c.CrawlerApplicationNumber(tx, appNumber)
				}(application.ApplicationNumber)
			}
			wg.Wait()

			pagination, err = c.storage.GetTaskApplicationNumberList(tx, taskId, currentPage+1, pageSize)
			if err != nil {
				collectLogger.Printf("[StartCrawler Task Id] %d Step 3 - Get Task Application Number (error: %s)", taskId, err)
				return fmt.Errorf("[StartCrawler Task Id] %d Step 3 - Get Task Application Number", taskId)
			}
		}

		if err := tx.Model(&currentTask).Update("completed", time.Now()).Error; err != nil {
			collectLogger.Printf("[StartCrawler Task Id] %d Step 4 - Update Task completed Date (error: %s)", taskId, err)
			return fmt.Errorf("[StartCrawler Task Id] %d Step 4 - Update Task completed Date", taskId)
		}
		return nil
	})
}

func (c *kiprisCollector) saveKiprisCollectorHistory(tx *gorm.DB, applicationNumber string, isSuccess bool, Error string) {
	history := model.KiprisCollectorHistory{
		ApplicationNumber: applicationNumber,
		IsSuccess:         isSuccess,
		Error:             Error,
	}
	err := tx.Create(&history)
	if err != nil {
		collectLogger.Printf("[ApplicationNumber] %s (error: %s)", applicationNumber, Error)
	}
}

func getKiprisTradeMarkInfo(c *kiprisCollector, applicationNumber string) (*model.KiprisResponse, string) {
	params := map[string]string{
		"applicationNumber": applicationNumber,
		"accessKey":         c.GetAccessKey(),
	}

	content, err := c.Get("/trademarkInfoSearchService/applicationNumberSearchInfo", params)
	if err != nil {
		return nil, fmt.Sprintf("[GET TradeMarkInfo %s] applicationNumberSearchInfo request step", applicationNumber)
	}

	var tradeMarkInfo model.KiprisResponse
	err = c.parser.Parse(content, &tradeMarkInfo)

	if err != nil {
		return nil, fmt.Sprintf("[GET TradeMarkInfo %s] applicationNumberSearchInfo parsing step", applicationNumber)
	}

	if tradeMarkInfo.Result() == model.Success {
		return &tradeMarkInfo, ""
	} else {
		return nil, fmt.Sprintf("[GET TradeMarkInfo %s] applicationNumberSearchInfo response %s", applicationNumber, getKiprisRequestResult(tradeMarkInfo.Result()))
	}
}

func getKiprisRequestResult(result model.KiprisResponseStatus) string {
	switch result {
	case model.Empty:
		return "Empty ApplicationNumber"
	case model.Error:
		return "Request Error"
	case model.Success:
		return "Request Success"
	default:
		return "None"
	}
}

func getTrademarkDesignationGoodstInfo(c *kiprisCollector, applicationNumber string) (*model.KiprisResponse, string) {
	params := map[string]string{
		"applicationNumber": applicationNumber,
		"accessKey":         c.GetAccessKey(),
	}

	content, err := c.Get("/trademarkInfoSearchService/trademarkDesignationGoodstInfo", params)
	if err != nil {
		return nil, fmt.Sprintf("[GET TrademarkDesignationGoodstInfo %s] trademarkDesignationGoodstInfo request step", applicationNumber)
	}

	var trademarkDesignationGoodstInfo model.KiprisResponse
	err = c.parser.Parse(content, &trademarkDesignationGoodstInfo)
	if err != nil {
		return nil, fmt.Sprintf("[GET TrademarkDesignationGoodstInfo %s] trademarkDesignationGoodstInfo pasring step", applicationNumber)
	}

	if trademarkDesignationGoodstInfo.Result() == model.Success {
		return &trademarkDesignationGoodstInfo, ""
	} else {
		return nil, fmt.Sprintf("[GET TrademarkDesignationGoodstInfo %s] trademarkDesignationGoodstInfo response %s", applicationNumber, getKiprisRequestResult(trademarkDesignationGoodstInfo.Result()))
	}
}

// TODO paralle
func (c *kiprisCollector) CrawlerApplicationNumber(tx *gorm.DB, applicationNumber string) bool {
	tradeMarkInfo, errString := getKiprisTradeMarkInfo(c, applicationNumber)

	if errString != "" {
		c.saveKiprisCollectorHistory(tx, applicationNumber, false, errString)
		return false
	}

	trademarkDesignationGoodstInfo, errString := getTrademarkDesignationGoodstInfo(c, applicationNumber)

	if errString != "" {
		c.saveKiprisCollectorHistory(tx, applicationNumber, false, errString)
		return false
	}

	if err := tx.Create(&tradeMarkInfo.Body.Items.TradeMarkInfo).Error; err != nil {
		c.saveKiprisCollectorHistory(tx, applicationNumber, false, err.Error())
		return false
	}

	for _, good := range trademarkDesignationGoodstInfo.Body.Items.TrademarkDesignationGoodstInfo {
		good.ApplicationNumber = applicationNumber
		if err := tx.Create(&good).Error; err != nil {
			c.saveKiprisCollectorHistory(tx, applicationNumber, false, err.Error())
			return false
		}
	}

	statistic := model.KiprisCollectorStatus{
		ApplicationNumber:                  applicationNumber,
		TradeMarkInfoStatus:                tradeMarkInfo.Result(),
		TradeMarkDesignationGoodInfoStatus: trademarkDesignationGoodstInfo.Result(),
	}

	if err := tx.Create(&statistic).Error; err != nil {
		c.saveKiprisCollectorHistory(tx, applicationNumber, false, err.Error())
		return false
	}

	c.saveKiprisCollectorHistory(tx, applicationNumber, true, "")

	return true
}

func getSerialNumberList(serialNumberRange string) []int {
	serialRangeList := strings.Split(serialNumberRange, ",")

	min := 0
	max := 9999999
	for _, serial := range serialRangeList {
		value, _ := strconv.Atoi(serial)
		if value >= min && value <= max {
			// validate range
			if min == 0 {
				min = value
			} else if value > min {
				max = value
			}
		} else {
			panic("Not validate serialNumber")
		}
	}

	serialNumberList := make([]int, max-min+1)
	for index, _ := range serialNumberList {
		serialNumberList[index] = index + min
	}

	return serialNumberList
}

func (c *kiprisCollector) CreateApplicationNumber(productCode string, year string, serialNumber int) string {
	result := fmt.Sprintf("%s%s%07d", productCode, year, serialNumber)
	return result
}

// func (c *kiprisCollector) GetYearLastApplicationNumber(year string) string {
// 	lastApplicationNumber := c.storage.GetYearLastApplicationNumber(year)
// 	return lastApplicationNumber
// }

// Not used
// func (c *kiprisCollector) GetMidValue(startNumber int, lastNumber int) int {
// 	// startNumber, lastNumber가 int 형이기 때문에
// 	// (lastNumber-startNumber)/2의 값은 버림처리가 된다.
// 	mid := (lastNumber-startNumber)/2 + startNumber
// 	return mid
// }

// func (c *kiprisCollector) GetLastApplicationNumber(startNumber string, lastNumber string, checker func(string) bool) (string, string, error) {
// 	start, err := strconv.Atoi(startNumber)
// 	last, err := strconv.Atoi(lastNumber)

// 	if err != nil {
// 		return "", "", err
// 	}

// 	if start >= last {
// 		return "", "", fmt.Errorf("uncorrect %d, %d", start, last)
// 	}

// 	mid := c.GetMidValue(start, last)

// 	midNumber := strconv.Itoa(mid)

// 	if mid == start {
// 		return startNumber, startNumber, nil
// 	} else if mid == last {
// 		return lastNumber, lastNumber, nil
// 	}

// 	isExist := checker(strconv.Itoa(mid))

// 	if isExist {
// 		return midNumber, lastNumber, nil
// 	}

// 	return startNumber, midNumber, nil
// }

// Not used
// func (c *kiprisCollector) IsTestApplicationNumberExist(answer string) func(string) bool {
// 	answerNumber, _ := strconv.Atoi(answer)
// 	return func(applicationNumber string) bool {
// 		number, _ := strconv.Atoi(applicationNumber)

// 		if number <= answerNumber {
// 			return true
// 		}
// 		return false
// 	}
// }

// Not used
// func (c *kiprisCollector) isApplicationNumberExist(applicationNumber string) bool {
// 	params := map[string]string{
// 		"applicationNumber": applicationNumber,
// 		"accessKey":         c.GetAccessKey(),
// 	}
// 	content, err := c.Get("/trademarkInfoSearchService/applicationNumberSearchInfo", params)

// 	if err != nil {
// 		return false
// 	}

// 	var tradeMarkInfo model.KiprisResponse
// 	c.parser.Parse(content, &tradeMarkInfo)
// 	if tradeMarkInfo.Result() == model.Success {
// 		return true
// 	}

// 	return false
// }

// select * from kipris_tasks;

// select * from kipris_application_numbers;

// select * from kipris_collector_histories
// where is_success = 0;

// select * from kipris_collector_histories
// where is_success = 1;

// select * from kipris_collector_statuses;

// select * from trade_mark_infos

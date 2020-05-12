package collector

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	// "sync"

	"github.com/gin-gonic/gin"
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

	// query, err := query.NewApp(types.QueryConfig{
	// 	DbType:       config.DbType,
	// 	DbConnString: config.DbConnString,
	// })

	if err != nil {
		return nil, err
	}

	return &kiprisCollector{
		endpt:     config.Endpoint,
		accessKey: config.AccessKey,
		parser:    parserInstance,
		storage:   storage,
		// query:     query,
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

func (c *kiprisCollector) getSearch(ctx *gin.Context) {
}

func (c *kiprisCollector) GetMethods() ([]types.RestMethod, error) {
	restMethods := make([]types.RestMethod, 1)
	restMethods = append(restMethods,
		types.RestMethod{
			Path: "/search",
			Handler: func(ctx *gin.Context) {
				param := types.TaskParameters{
					ProductCode:       "40",
					Year:              "2017",
					SerialNumberRange: "1,20",
				}
				c.GetApplicationNumberList(param)
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
				params := types.TaskParameters{
					ProductCode:       "40",
					Year:              "2017",
					SerialNumberRange: "1,20",
				}
				pagination, err := c.GetApplicationNumberList(params)
				fmt.Println(pagination.Data)
				fmt.Println(err)
			},
		},
	)
	return restMethods, nil
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

// create kipris collector task
// 처음에는 create만 만들고 추후에 delete & create 하자
func (c *kiprisCollector) CreatTask(args types.TaskParameters) error {
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

	// kiprisApplicationNumbers := make([]model.KiprisApplicationNumber, 0)

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
			// kiprisApplicationNumbers = append(kiprisApplicationNumbers, kiprisApplicationNumber)
			err := c.storage.Create(&kiprisApplicationNumber)
			if err != nil {
				return err
			}
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
				// kiprisApplicationNumbers = append(kiprisApplicationNumbers, kiprisApplicationNumber)
				err := c.storage.Create(&kiprisApplicationNumber)
				if err != nil {
					return err
				}
			}
		}
	}

	// fmt.Println(kiprisApplicationNumbers)
	return nil
}

func (c *kiprisCollector) CreatManualTask(args types.TaskParameters) error {
	productCode := args.ProductCode
	year := args.Year
	serialNumberRange := args.SerialNumberRange

	// require
	if productCode == "" || year == "" || serialNumberRange == "" {
		return fmt.Errorf("Please pass task parameters")
	}

	serialNumberList := getSerialNumberList(serialNumberRange)
	for _, serialNumber := range serialNumberList {
		kiprisApplicationNumber := model.KiprisApplicationNumber{
			ApplicationNumber: fmt.Sprintf("%s%s%07d", productCode, year, serialNumber),
			ProductCode:       "40",
			Year:              year,
			SerialNumber:      serialNumber,
		}
		err := c.storage.Create(&kiprisApplicationNumber)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *kiprisCollector) GetApplicationNumberList(args types.TaskParameters) (*pagination.Paginator, error) {
	searchResult := make([]model.KiprisApplicationNumber, 0)
	searchData := model.KiprisApplicationNumber{
		Year:        args.Year,
		ProductCode: args.ProductCode,
	}

	startSerialNumber := 0
	endSerialNumber := 0
	serialRangeList := strings.Split(args.SerialNumberRange, ",")
	if len(serialRangeList) == 2 {
		startSerialNumber, _ = strconv.Atoi(serialRangeList[0])
		endSerialNumber, _ = strconv.Atoi(serialRangeList[1])
	}

	pagination, err := c.storage.GetKiprisApplicationNumberList(searchData, &searchResult, startSerialNumber, endSerialNumber, args.Page, args.Size)
	return pagination, err
}

// product string, year string, length int, startNumber int
func (c *kiprisCollector) StartCrawler(year string, productCode string, startSerialNumber int, endSerialNumber int) {
	// searchResult := make([]model.KiprisApplicationNumber, 0)
	// searchData := model.KiprisApplicationNumber{
	// 	Year:        year,
	// 	ProductCode: productCode,
	// }

	// c.storage.GetKiprisApplicationNumberList(searchData, &searchResult, startSerialNumber, endSerialNumber)

	// // Crawer history

	// var wg sync.WaitGroup
	// for _, application := range searchResult {
	// 	wg.Add(1)
	// 	go func(appNumber string) {
	// 		defer wg.Done()
	// 		c.CrawlerApplicationNumber(appNumber)
	// 	}(application.ApplicationNumber)
	// }
	// wg.Wait()
}

func (c *kiprisCollector) saveKiprisCollectorHistory(applicationNumber string, isSuccess bool, Error string) {
	history := model.KiprisCollectorHistory{
		ApplicationNumber: applicationNumber,
		IsSuccess:         isSuccess,
		Error:             Error,
	}
	if Error == "" {
		log.Printf("[Success ApplicationNumber] %s", applicationNumber)
	} else {
		log.Printf("[Fail ApplicationNumber] %s (error: %s)", applicationNumber, Error)
	}
	err := c.storage.Create(&history)
	if err != nil {
		log.Printf("[Save History ApplicationNumber] %s (error: %s)", applicationNumber, Error)
		// collectLogger.Printf("[ApplicationNumber] %s (error: %s)", applicationNumber, Error)
	}
}

func getKiprisTradeMarkInfo(c *kiprisCollector, applicationNumber string) (*model.KiprisResponse, string) {
	params := map[string]string{
		"applicationNumber": applicationNumber,
		"accessKey":         c.GetAccessKey(),
	}
	content, err := c.Get("/trademarkInfoSearchService/applicationNumberSearchInfo", params)
	if err != nil {
		return nil, fmt.Sprint("[GET TradeMarkInfo] applicationNumberSearchInfo request step")
	}

	var tradeMarkInfo model.KiprisResponse
	err = c.parser.Parse(content, &tradeMarkInfo)

	if err != nil {
		return nil, fmt.Sprint("[GET TradeMarkInfo] applicationNumberSearchInfo parsing step")
	}

	if tradeMarkInfo.Result() == model.Success {
		return &tradeMarkInfo, ""
	} else {
		return nil, fmt.Sprintf("[GET TradeMarkInfo] applicationNumberSearchInfo response %d", tradeMarkInfo.Result())
	}
}

func getTrademarkDesignationGoodstInfo(c *kiprisCollector, applicationNumber string) (*model.KiprisResponse, string) {
	params := map[string]string{
		"applicationNumber": applicationNumber,
		"accessKey":         c.GetAccessKey(),
	}

	content, err := c.Get("/trademarkInfoSearchService/trademarkDesignationGoodstInfo", params)
	if err != nil {
		return nil, fmt.Sprint("[GET TrademarkDesignationGoodstInfo] trademarkDesignationGoodstInfo request step")
	}

	var trademarkDesignationGoodstInfo model.KiprisResponse
	err = c.parser.Parse(content, &trademarkDesignationGoodstInfo)
	if err != nil {
		return nil, fmt.Sprint("[GET TrademarkDesignationGoodstInfo] trademarkDesignationGoodstInfo pasring step")
	}

	if trademarkDesignationGoodstInfo.Result() == model.Success {
		return &trademarkDesignationGoodstInfo, ""
	} else {
		return nil, fmt.Sprintf("[GET TrademarkDesignationGoodstInfo] trademarkDesignationGoodstInfo response %d", trademarkDesignationGoodstInfo.Result())
	}
}

// 병렬처리
func (c *kiprisCollector) CrawlerApplicationNumber(applicationNumber string) bool {
	tradeMarkInfo, errString := getKiprisTradeMarkInfo(c, applicationNumber)

	if errString != "" {
		c.saveKiprisCollectorHistory(applicationNumber, false, errString)
		return false
	}

	trademarkDesignationGoodstInfo, errString := getTrademarkDesignationGoodstInfo(c, applicationNumber)

	if errString != "" {
		c.saveKiprisCollectorHistory(applicationNumber, false, errString)
		return false
	}

	err := c.storage.Create(&tradeMarkInfo.Body.Items.TradeMarkInfo)
	if err != nil {
		c.saveKiprisCollectorHistory(applicationNumber, false, err.Error())
		return false
	}

	for _, good := range trademarkDesignationGoodstInfo.Body.Items.TrademarkDesignationGoodstInfo {
		good.ApplicationNumber = applicationNumber
		err := c.storage.Create(&good)
		if err != nil {
			c.saveKiprisCollectorHistory(applicationNumber, false, err.Error())
			return false
		}
	}

	statistic := model.KiprisCollectorStatus{
		ApplicationNumber:                  applicationNumber,
		TradeMarkInfoStatus:                tradeMarkInfo.Result(),
		TradeMarkDesignationGoodInfoStatus: trademarkDesignationGoodstInfo.Result(),
	}

	err = c.storage.Create(&statistic)

	if err != nil {
		c.saveKiprisCollectorHistory(applicationNumber, false, err.Error())
		return false
	}

	c.saveKiprisCollectorHistory(applicationNumber, true, "")

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

func (c *kiprisCollector) CreateApplicationNumberList(year string, length int, startNumber int) []string {
	// 상표법 개정에 따라 서비스표(41), 상표/서비스표(45)는 2016년 9월 1일 이후 출원건에 대해 상표(40)에 통합 되었습니다.
	applicationNumberList := make([]string, 0)

	// productCodeList := []string{
	// 	"40",
	// 	"41",
	// 	"45",
	// }

	// var value model.KiprisApplicationNumber

	// yearNum, _ := strconv.Atoi(year)
	// if yearNum > 2016 {
	// 	//  40
	// 	serialNumberList := getSerialNumberList("40", year, length, startNumber)
	// 	for _, serialNumber := range serialNumberList {
	// 		value = model.KiprisApplicationNumber{
	// 			ApplicationNumber: fmt.Sprintf("%s%s%07d", "40", year, serialNumber),
	// 			ProductCode:       "40",
	// 			Year:              year,
	// 			SerialNumber:      serialNumber,
	// 		}
	// 		c.storage.Create(&value)
	// 	}
	// } else {
	// 	for _, productCode := range productCodeList {
	// 		serialNumberList := getSerialNumberList(productCode, year, length, startNumber)
	// 		for _, serialNumber := range serialNumberList {
	// 			value = model.KiprisApplicationNumber{
	// 				ApplicationNumber: fmt.Sprintf("%s%s%07d", productCode, year, serialNumber),
	// 				ProductCode:       productCode,
	// 				Year:              year,
	// 				SerialNumber:      serialNumber,
	// 			}
	// 			c.storage.Create(&value)
	// 		}
	// 	}
	// }

	return applicationNumberList
}

func (c *kiprisCollector) CreateApplicationNumber(productCode string, year string, serialNumber int) string {
	result := fmt.Sprintf("%s%s%07d", productCode, year, serialNumber)
	return result
}

func (c *kiprisCollector) GetYearLastApplicationNumber(year string) string {
	lastApplicationNumber := c.storage.GetYearLastApplicationNumber(year)
	return lastApplicationNumber
}

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

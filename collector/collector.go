package collector

import (
	"fmt"
	"kipris-collector/model"
	"kipris-collector/parser"
	"kipris-collector/storage"
	"kipris-collector/types"
	"kipris-collector/utils"
	"log"
	"os"
	"strconv"
)

type kiprisCollector struct {
	endpt     string
	accessKey string
	parser    types.Parser
	storage   types.Storage
}

var collectLogger *log.Logger

func init() {
	fpLog, err := os.OpenFile("collector_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// defer fpLog.Close()

	collectLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewCollector(config collectorConfig) (types.Collector, error) {
	parserInstance, err := parser.NewParser("xml")

	if err != nil {
		return nil, err
	}

	storage, err := storage.NewStorage(storage.StorageConfig{
		DbType:       config.DbType,
		DbConnString: config.DbConnString,
	})

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

func (c *kiprisCollector) CrawlerApplicationNumber(applicationNumber string) bool {
	// collectLogger.Printf("applicationNumber: %s start", applicationNumber)
	params := map[string]string{
		"applicationNumber": applicationNumber,
		"accessKey":         c.GetAccessKey(),
	}
	content, err := c.Get("/trademarkInfoSearchService/applicationNumberSearchInfo", params)

	if err != nil {
		collectLogger.Printf("applicationNumber: %s error %s", applicationNumber, err)
		return false
	}

	var tradeMarkInfo model.KiprisResponse
	err = c.parser.Parse(content, &tradeMarkInfo)
	if err != nil {
		collectLogger.Printf("applicationNumber: %s error %s", applicationNumber, err)
		return false
	}
	err = c.storage.Create(&tradeMarkInfo.Body.Items.TradeMarkInfo)
	if err != nil {
		collectLogger.Printf("applicationNumber: %s error %s", applicationNumber, err)
		return false
	}

	content, err = c.Get("/trademarkInfoSearchService/trademarkDesignationGoodstInfo", params)
	if err != nil {
		collectLogger.Printf("applicationNumber: %s error %s", applicationNumber, err)
		return false
	}

	var trademarkDesignationGoodstInfo model.KiprisResponse
	err = c.parser.Parse(content, &trademarkDesignationGoodstInfo)
	if err != nil {
		collectLogger.Printf("applicationNumber: %s error %s", applicationNumber, err)
		return false
	}

	for _, good := range trademarkDesignationGoodstInfo.Body.Items.TrademarkDesignationGoodstInfo {
		good.ApplicationNumber = applicationNumber
		err := c.storage.Create(&good)
		if err != nil {
			collectLogger.Printf("applicationNumber: %s error %s", applicationNumber, err)
			return false
		}
	}

	statistic := model.KiprisCollector{
		ApplicationNumber:                  applicationNumber,
		TradeMarkInfoStatus:                tradeMarkInfo.Result(),
		TradeMarkDesignationGoodInfoStatus: trademarkDesignationGoodstInfo.Result(),
	}

	err = c.storage.Create(&statistic)

	if err != nil {
		collectLogger.Printf("applicationNumber: %s error %s", applicationNumber, err)
		return false
	}

	// collectLogger.Printf("applicationNumber: %s end", applicationNumber)
	return true
}

func (c *kiprisCollector) isApplicationNumberExist(applicationNumber string) bool {
	params := map[string]string{
		"applicationNumber": applicationNumber,
		"accessKey":         c.GetAccessKey(),
	}
	content, err := c.Get("/trademarkInfoSearchService/applicationNumberSearchInfo", params)

	if err != nil {
		return false
	}

	var tradeMarkInfo model.KiprisResponse
	c.parser.Parse(content, &tradeMarkInfo)
	if tradeMarkInfo.Result() == model.Success {
		return true
	}

	return false
}

func (c *kiprisCollector) CreateApplicationNumberList() []string {
	// 상표법 개정에 따라 서비스표(41), 상표/서비스표(45)는 2016년 9월 1일 이후 출원건에 대해 상표(40)에 통합 되었습니다.

	applicationNumberList := make([]string, 0)

	// productCodeList := []string{
	// 	"40",
	// 	"41",
	// 	"45",
	// }

	yearList := make([]string, 0)
	current, _ := strconv.Atoi("2020")
	start, _ := strconv.Atoi("1950")

	for i := start; i <= current; i++ {
		yearList = append(yearList, strconv.Itoa(i))
	}

	serialNumberList := make([]string, 100)
	for index, _ := range serialNumberList {
		serialNumberList[index] = fmt.Sprintf("%07d", index+1)
		applicationNumberList = append(applicationNumberList, fmt.Sprintf("%s%s%s", "40", "2020", serialNumberList[index]))
	}

	// for _, year := range yearList {
	// 	yearNum, _ := strconv.Atoi(year)
	// 	if yearNum > 2016 {
	// 		//  40
	// 	} else {
	// 		// productCodeList
	// 	}
	// }

	return applicationNumberList
}

func (c *kiprisCollector) CreateApplicationNumber(productCode string, year string, serialNumber int) string {
	result := fmt.Sprintf("%s%s%07d", productCode, year, serialNumber)
	return result
}

func (c *kiprisCollector) GetMidValue(startNumber int, lastNumber int) int {
	// startNumber, lastNumber가 int 형이기 때문에
	// (lastNumber-startNumber)/2의 값은 버림처리가 된다.
	mid := (lastNumber-startNumber)/2 + startNumber
	return mid
}

func (c *kiprisCollector) GetLastApplicationNumber(startNumber string, lastNumber string, checker func(string) bool) (string, string, error) {
	start, err := strconv.Atoi(startNumber)
	last, err := strconv.Atoi(lastNumber)

	if err != nil {
		return "", "", err
	}

	if start >= last {
		return "", "", fmt.Errorf("uncorrect %d, %d", start, last)
	}

	mid := c.GetMidValue(start, last)

	midNumber := strconv.Itoa(mid)

	if mid == start {
		return startNumber, startNumber, nil
	} else if mid == last {
		return lastNumber, lastNumber, nil
	}

	isExist := checker(strconv.Itoa(mid))

	if isExist {
		return midNumber, lastNumber, nil
	}

	return startNumber, midNumber, nil
}

// for test
func (c *kiprisCollector) IsTestApplicationNumberExist(answer string) func(string) bool {
	answerNumber, _ := strconv.Atoi(answer)
	return func(applicationNumber string) bool {
		number, _ := strconv.Atoi(applicationNumber)

		if number <= answerNumber {
			return true
		}
		return false
	}
}

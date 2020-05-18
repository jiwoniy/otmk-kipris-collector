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
	"github.com/jiwoniy/otmk-kipris-collector/kipris/model"
	"github.com/jiwoniy/otmk-kipris-collector/kipris/parser"
	"github.com/jiwoniy/otmk-kipris-collector/kipris/storage"
	"github.com/jiwoniy/otmk-kipris-collector/kipris/types"
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
var crawlSize = 25

func init() {
	fpLog, err := os.OpenFile("collector.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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

func writeError(c *gin.Context, err error) {
	response := types.RestFailResponse{
		Error: err.Error(),
	}
	c.JSON(http.StatusBadRequest, response)
}

// For rest
func (c *kiprisCollector) GetMethods() ([]types.RestMethod, error) {
	restMethods := make([]types.RestMethod, 1)
	restMethods = append(restMethods,
		types.RestMethod{
			Path: "/ping",
			Handler: func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "pong")
			},
		},
		types.RestMethod{
			Path: "/tasks",
			Handler: func(ctx *gin.Context) {
				page := ctx.DefaultQuery("page", "1")
				size := ctx.DefaultQuery("50", "50")

				pagination, err := c.GetTaskList(page, size)

				if err != nil {
					writeError(ctx, err)
				} else {
					ctx.JSON(http.StatusOK, pagination)
				}
			},
		},
		types.RestMethod{
			Path: "/task/:taskId",
			Handler: func(ctx *gin.Context) {
				taskId := ctx.Param("taskId")
				uTaskId, _ := strconv.ParseInt(taskId, 10, 64)
				data, _ := c.GetTaskById(uTaskId)
				if data.ID == 0 {
					ctx.String(http.StatusOK, "not exists taskId")
				} else {
					ctx.JSON(http.StatusOK, data)
				}
			},
		},
	)
	return restMethods, nil
}

func (c *kiprisCollector) PostMethods() ([]types.RestMethod, error) {
	restMethods := make([]types.RestMethod, 1)
	restMethods = append(restMethods,
		types.RestMethod{
			Path: "/task",
			Handler: func(ctx *gin.Context) {
				var param types.TaskParameters
				if err := ctx.ShouldBind(&param); err != nil {
					writeError(ctx, err)
					return
				}

				err := c.CreateTask(param)
				if err != nil {
					writeError(ctx, err)
					return
				} else {
					ctx.String(http.StatusOK, "create task success")
					return
				}
			},
		},
		types.RestMethod{
			Path: "/task/:taskId",
			Handler: func(ctx *gin.Context) {
				taskId := ctx.Param("taskId")

				uTaskId, _ := strconv.ParseInt(taskId, 10, 64)
				data, _ := c.GetTaskById(uTaskId)
				if data.ID == 0 {
					ctx.String(http.StatusOK, "not exists taskId")
					return
				} else {
					go func() {
						c.StartCrawler(uTaskId)
					}()
					ctx.String(http.StatusOK, fmt.Sprintf("start task %s", taskId))

				}
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
		yearLastApplicationSerialNumber := c.storage.GetYearLastApplicationSerialNumber(year)
		if yearLastApplicationSerialNumber == 0 {
			serialNumberRange = "1," + strconv.Itoa(taskNumber)
		} else {
			start := yearLastApplicationSerialNumber
			newStart := start + 1
			newEnd := start + taskNumber
			serialNumberRange = strconv.Itoa(newStart) + "," + strconv.Itoa(newEnd)
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

func (c *kiprisCollector) GetTaskList(pageParam string, sizeParam string) (*utils.Paginator, error) {
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return nil, err
	}
	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		return nil, err
	}
	pagination, err := c.storage.GetTaskList(page, size)
	return pagination, err
}

func (c *kiprisCollector) GetTaskById(taskId int64) (model.KiprisTask, error) {
	task, err := c.storage.GetTaskById(taskId)
	if err != nil {
		return model.KiprisTask{}, err
	}
	return task, nil
}

func (c *kiprisCollector) StartCrawler(taskId int64) error {
	db := c.storage.GetDB()
	pageSize := crawlSize

	currentTask := model.KiprisTask{}

	if err := db.Table("kipris_tasks").Where("id = ?", taskId).First(&currentTask).Error; err != nil {
		collectLogger.Printf("[StartCrawler Task Id] %d Step 1 - Get Task (error: %s)", taskId, err)
		return fmt.Errorf("[StartCrawler Task Id] %d Step 1 - Get Task", taskId)
	}

	if err := db.Model(&currentTask).Update("startedAt", time.Now()).Error; err != nil {
		collectLogger.Printf("[StartCrawler Task Id] %d Step 2 - Update Task started Date (error: %s)", taskId, err)
		return fmt.Errorf("[StartCrawler Task Id] %d Step 2 - Update Task started Date", taskId)
	}

	paginationOut, err := c.storage.GetTaskApplicationNumberList(db, taskId, 1, pageSize)
	if err != nil {
		collectLogger.Printf("[StartCrawler Task Id] %d Step 3 - Get Task Application Number (error: %s)", taskId, err)
		return fmt.Errorf("[StartCrawler Task Id] %d Step 3 - Get Task Application Number", taskId)
	}

	return db.Transaction(func(tx *gorm.DB) error {
		totalPage := paginationOut.TotalPage
		currentPage := paginationOut.Page
		data := paginationOut.Data.(*[]model.KiprisApplicationNumber)

		var wgOut sync.WaitGroup
		for ; currentPage <= totalPage; currentPage++ {
			if currentPage > 1 {
				paginationIn, err := c.storage.GetTaskApplicationNumberList(db, taskId, currentPage, pageSize)
				if err != nil {
					collectLogger.Printf("[StartCrawler Task Id] %d Step 3 - Get Task Application Number (error: %s)", taskId, err)
					return fmt.Errorf("[StartCrawler Task Id] %d Step 3 - Get Task Application Number", taskId)
				}
				data = paginationIn.Data.(*[]model.KiprisApplicationNumber)
			}

			wgOut.Wait()
			wgOut.Add(1)

			var wg sync.WaitGroup
			for _, application := range *data {
				wg.Add(1)
				go func(appNumber string) {
					c.CrawlerApplicationNumber(tx, taskId, appNumber)
					defer wg.Done()
				}(application.ApplicationNumber)
			}
			wg.Wait()
			wgOut.Done()

		}

		if err := tx.Model(&currentTask).Update("completedAt", time.Now()).Error; err != nil {
			collectLogger.Printf("[StartCrawler Task Id] %d Step 4 - Update Task completed Date (error: %s)", taskId, err)
			return fmt.Errorf("[StartCrawler Task Id] %d Step 4 - Update Task completed Date", taskId)
		}
		return nil
	})
}

func (c *kiprisCollector) saveKiprisCollectorHistory(tx *gorm.DB, taskId int64, applicationNumber string, isSuccess bool, Error string) {
	history := model.KiprisCollectorHistory{
		ApplicationNumber: applicationNumber,
		IsSuccess:         isSuccess,
		Error:             Error,
		TaskId:            taskId,
	}
	err := tx.Create(&history).Error
	if err != nil {
		collectLogger.Printf("[SaveKiprisCollectorHistory] %s (error: %s)", applicationNumber, err)
	}
}

func getKiprisTradeMarkInfo(c *kiprisCollector, applicationNumber string) (*model.KiprisResponse, string) {
	params := map[string]string{
		"applicationNumber": applicationNumber,
		"accessKey":         c.GetAccessKey(),
	}

	content, err := c.Get("/trademarkInfoSearchService/applicationNumberSearchInfo", params)
	if err != nil {
		return nil, fmt.Sprintf("[TradeMarkInfo] applicationNumberSearchInfo request (error: %s)", err)
	}

	var tradeMarkInfo model.KiprisResponse
	err = c.parser.Parse(content, &tradeMarkInfo)

	if err != nil {
		return nil, fmt.Sprintf("[TradeMarkInfo] applicationNumberSearchInfo parsing (error: %s)", err)
	}

	if tradeMarkInfo.Result() == model.Success {
		return &tradeMarkInfo, ""
	} else {
		collectLogger.Printf("[TradeMarkInfo: %s] applicationNumberSearchInfo (response: %v)", applicationNumber, tradeMarkInfo.Body)
		return nil, fmt.Sprintf("[TradeMarkInfo] applicationNumberSearchInfo response %s", getKiprisRequestResult(tradeMarkInfo.Result()))
	}
}

func getKiprisRequestResult(result model.KiprisResponseStatus) string {
	switch result {
	case model.Empty:
		return "empty"
	case model.Error:
		return "error"
	case model.Success:
		return "success"
	default:
		return "none"
	}
}

func getTrademarkDesignationGoodstInfo(c *kiprisCollector, applicationNumber string) (*model.KiprisResponse, string) {
	params := map[string]string{
		"applicationNumber": applicationNumber,
		"accessKey":         c.GetAccessKey(),
	}

	content, err := c.Get("/trademarkInfoSearchService/trademarkDesignationGoodstInfo", params)
	if err != nil {
		return nil, fmt.Sprintf("[TrademarkDesignationGoodstInfo] trademarkDesignationGoodstInfo request (error: %s)", err)
	}

	var trademarkDesignationGoodstInfo model.KiprisResponse
	err = c.parser.Parse(content, &trademarkDesignationGoodstInfo)
	if err != nil {
		return nil, fmt.Sprintf("[TrademarkDesignationGoodstInfo] trademarkDesignationGoodstInfo pasring (error: %s)", err)
	}

	if trademarkDesignationGoodstInfo.Result() == model.Success {
		return &trademarkDesignationGoodstInfo, ""
	} else {
		collectLogger.Printf("[TrademarkDesignationGoodstInfo: %s] applicationNumberSearchInfo (response: %v)", applicationNumber, trademarkDesignationGoodstInfo.Body)
		return nil, fmt.Sprintf("[TrademarkDesignationGoodstInfo] trademarkDesignationGoodstInfo response %s (error: %s)", getKiprisRequestResult(trademarkDesignationGoodstInfo.Result()), err)
	}
}

func (c *kiprisCollector) CrawlerApplicationNumber(tx *gorm.DB, taskId int64, applicationNumber string) bool {

	// tradeMarkInfo, trademarkDesignationGoodstInfo, errString := c.getData(applicationNumber)

	// if errString != "" {
	// 	c.saveKiprisCollectorHistory(tx, taskId, applicationNumber, false, errString)
	// 	return false
	// }

	// isSuccess := c.saveData(tx, taskId, applicationNumber, tradeMarkInfo, trademarkDesignationGoodstInfo)

	// return isSuccess

	tradeMarkInfo, errString := getKiprisTradeMarkInfo(c, applicationNumber)
	if errString != "" {
		c.saveKiprisCollectorHistory(tx, taskId, applicationNumber, false, errString)
		return false
	}

	trademarkDesignationGoodstInfo, errString := getTrademarkDesignationGoodstInfo(c, applicationNumber)

	if errString != "" {
		c.saveKiprisCollectorHistory(tx, taskId, applicationNumber, false, errString)
		return false
	}

	if err := tx.Create(&tradeMarkInfo.Body.Items.TradeMarkInfo).Error; err != nil {
		c.saveKiprisCollectorHistory(tx, taskId, applicationNumber, false, err.Error())
		return false
	}

	for _, good := range trademarkDesignationGoodstInfo.Body.Items.TrademarkDesignationGoodstInfo {
		good.ApplicationNumber = applicationNumber
		if err := tx.Create(&good).Error; err != nil {
			c.saveKiprisCollectorHistory(tx, taskId, applicationNumber, false, err.Error())
			return false
		}
	}

	statistic := model.KiprisCollectorStatus{
		ApplicationNumber:                  applicationNumber,
		TaskId:                             taskId,
		TradeMarkInfoStatus:                tradeMarkInfo.Result(),
		TradeMarkDesignationGoodInfoStatus: trademarkDesignationGoodstInfo.Result(),
	}

	if err := tx.Create(&statistic).Error; err != nil {
		c.saveKiprisCollectorHistory(tx, taskId, applicationNumber, false, err.Error())
		return false
	}

	c.saveKiprisCollectorHistory(tx, taskId, applicationNumber, true, "")
	return true
}

func (c *kiprisCollector) getData(applicationNumber string) (*model.KiprisResponse, *model.KiprisResponse, string) {
	tradeMarkInfo, errString := getKiprisTradeMarkInfo(c, applicationNumber)
	if errString != "" {
		return nil, nil, errString
	}

	trademarkDesignationGoodstInfo, errString := getTrademarkDesignationGoodstInfo(c, applicationNumber)

	if errString != "" {
		return tradeMarkInfo, nil, errString
	}

	return tradeMarkInfo, trademarkDesignationGoodstInfo, ""
}

func (c *kiprisCollector) saveData(tx *gorm.DB, taskId int64, applicationNumber string, tradeMarkInfo *model.KiprisResponse, trademarkDesignationGoodstInfo *model.KiprisResponse) bool {
	if err := tx.Create(&tradeMarkInfo.Body.Items.TradeMarkInfo).Error; err != nil {
		c.saveKiprisCollectorHistory(tx, taskId, applicationNumber, false, err.Error())
		return false
	}

	for _, good := range trademarkDesignationGoodstInfo.Body.Items.TrademarkDesignationGoodstInfo {
		good.ApplicationNumber = applicationNumber
		if err := tx.Create(&good).Error; err != nil {
			c.saveKiprisCollectorHistory(tx, taskId, applicationNumber, false, err.Error())
			return false
		}
	}

	statistic := model.KiprisCollectorStatus{
		ApplicationNumber:                  applicationNumber,
		TaskId:                             taskId,
		TradeMarkInfoStatus:                tradeMarkInfo.Result(),
		TradeMarkDesignationGoodInfoStatus: trademarkDesignationGoodstInfo.Result(),
	}

	if err := tx.Create(&statistic).Error; err != nil {
		c.saveKiprisCollectorHistory(tx, taskId, applicationNumber, false, err.Error())
		return false
	}

	c.saveKiprisCollectorHistory(tx, taskId, applicationNumber, true, "")
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

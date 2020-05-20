package collector

import (
	"testing"

	"github.com/jiwoniy/otmk-kipris-collector/kipris/types"

	"github.com/stretchr/testify/suite"
)

type CollectorTestSuite struct {
	suite.Suite
	collector types.Collector
}

func (suite *CollectorTestSuite) SetupTest() {
	config := types.CollectorConfig{
		KiprisConfig: types.KiprisConfig{
			Endpoint:  "http://plus.kipris.or.kr/openapi/rest",
			AccessKey: "I0Jnw4w6/UpQSp1zHPsIDSztV9=hgVUNI6IANH3bCEw=", // onthe mark key
		},
		DbConfig: types.DbConfig{
			DbType:       "sqlite3",
			DbConnString: ":memory:",
			// DbType:       "mysql",
			// DbConnString: "kipris_admin:Kipris0518!@@(61.97.187.142:3306)/kipris_collector?charset=utf8&parseTime=True&loc=Local",
		},
		// DbType:       "mysql",
		// DbConnString: "kipris_server:OnthemarkKipris0507!@@(61.97.187.142:3306)/kipris?charset=utf8&parseTime=True&loc=Local",
		// DbConnString: "./test.db",
	}

	collector, err := NewCollector(config)
	if err != nil {
		suite.Error(err)
	}
	suite.collector = collector
}

// func (suite *CollectorTestSuite) TestCollector() {
// 	parserInstance := suite.collector.GetParser()

// 	type testcases struct {
// 		url            string
// 		params         map[string]string
// 		dest           model.KiprisResponse
// 		responseStatus model.KiprisResponseStatus
// 	}
// 	tests := []testcases{
// 		{
// 			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
// 			params: map[string]string{
// 				"applicationNumber": "4020200023099", // valid number
// 				"accessKey":         suite.collector.GetAccessKey(),
// 			},
// 			dest:           model.KiprisResponse{},
// 			responseStatus: model.Success,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
// 			params: map[string]string{
// 				"applicationNumber": "402020002309911", // invalid number
// 				"accessKey":         suite.collector.GetAccessKey(),
// 			},
// 			dest:           model.KiprisResponse{},
// 			responseStatus: model.Empty,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
// 			params: map[string]string{
// 				"applicationNumber": "", // invalid number
// 				"accessKey":         suite.collector.GetAccessKey(),
// 			},
// 			dest:           model.KiprisResponse{},
// 			responseStatus: model.Error,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
// 			params: map[string]string{
// 				"applicationNumber": "4020200023099", // valid number
// 				"accessKey":         "",
// 			},
// 			dest:           model.KiprisResponse{},
// 			responseStatus: model.Error,
// 		},

// 		{
// 			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
// 			params: map[string]string{
// 				"applicationNumber": "4020200023099", // valid number
// 				"accessKey":         suite.collector.GetAccessKey(),
// 			},
// 			dest:           model.KiprisResponse{},
// 			responseStatus: model.Success,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
// 			params: map[string]string{
// 				"applicationNumber": "402020002309911", // invalid number
// 				"accessKey":         suite.collector.GetAccessKey(),
// 			},
// 			dest:           model.KiprisResponse{},
// 			responseStatus: model.Empty,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
// 			params: map[string]string{
// 				"applicationNumber": "", // invalid number
// 				"accessKey":         suite.collector.GetAccessKey(),
// 			},
// 			dest:           model.KiprisResponse{},
// 			responseStatus: model.Error,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
// 			params: map[string]string{
// 				"applicationNumber": "4020200023099", // valid number
// 				"accessKey":         "",
// 			},
// 			dest:           model.KiprisResponse{},
// 			responseStatus: model.Error,
// 		},
// 	}

// 	for testIndex, tc := range tests {
// 		content, err := suite.collector.Get(tc.url, tc.params)
// 		if err != nil {
// 			suite.Error(err)
// 		}

// 		parserInstance.Parse(content, &tc.dest)

// 		if tc.dest.Result() != tc.responseStatus {
// 			suite.Error(fmt.Errorf("This test index fail %d", testIndex))
// 		}
// 	}
// }

// func (suite *CollectorTestSuite) TestGetYearLastApplicationNumber() {
// 	lastApplicationNumber := suite.collector.GetYearLastApplicationNumber("2017")
// 	fmt.Println(lastApplicationNumber)
// }

// func (suite *CollectorTestSuite) TestCreaManulTask() {
// 	param := types.TaskParameters{
// 		ProductCode:       "40",
// 		Year:              "2015",
// 		SerialNumberRange: "2,3",
// 	}
// 	err := suite.collector.CreatManualTask(param)
// 	fmt.Println(err)
// }

// func (suite *CollectorTestSuite) TestGetTaskList() {
// 	pagination, err := suite.collector.GetTaskList(1, 20)
// }

// func (suite *CollectorTestSuite) TestGetTaskApplicationNumberList() {
// 	pagination, _ := suite.collector.GetTaskApplicationNumberList(1, 1, 20)
// 	fmt.Println(pagination.TotalRecord)
// }

// func (suite *CollectorTestSuite) TestStartCrawler() {
// 	param := types.TaskParameters{
// 		Year:              "2018",
// 		SerialNumberRange: "1,10",
// 	}
// 	suite.collector.CreateTask(param)
// 	suite.collector.StartCrawler(1)
// }

// func (suite *CollectorTestSuite) TestStarAutotCrawler() {
// 	suite.collector.StartAutoCrawler()
// }

// func (suite *CollectorTestSuite) TestGetKiprisApplicationNumerList() {
// 	param := types.TaskParameters{
// 		ProductCode:       "40",
// 		Year:              "2017",
// 		SerialNumberRange: "1,20",
// 	}
// 	pagination, err := suite.collector.GetApplicationNumberList(param)
// 	fmt.Println(pagination.Data)
// 	fmt.Println(err)
// }

// func (suite *CollectorTestSuite) TestCrawler() {
// 	suite.collector.StartCrawler("2017", "40", 300, 350)
// }

// func (suite *CollectorTestSuite) TestCreateApplicationNumberList() {
// 	suite.collector.CreateApplicationNumberList("2017", 1000, 1)
// }

// func (suite *CollectorTestSuite) TestTotalCreateApplicationNumberList() {
// 	yearList := make([]string, 0)
// 	current, _ := strconv.Atoi("2020")
// 	start, _ := strconv.Atoi("1950")

// 	for i := start; i <= current; i++ {
// 		yearList = append(yearList, strconv.Itoa(i))
// 	}

// 	for _, year := range yearList {
// 		applicationNumberList := suite.collector.CreateApplicationNumberList(year, 100)
// 		fmt.Println(applicationNumberList)
// 	}
// }

// func (suite *CollectorTestSuite) TestGetApplicationNumberList() {
// 	suite.collector.CreateApplicationNumberList("1998", 100, 200)
// }

// Not used
// func (suite *CollectorTestSuite) TestFindApplicationNumberLogic() {
// 	type testcases struct {
// 		startNumber string
// 		lastNumber  string
// 		findNumber  string
// 	}
// 	tests := []testcases{
// 		{
// 			startNumber: "1",
// 			lastNumber:  "300",
// 			findNumber:  "275",
// 		},
// 		{
// 			startNumber: "1",
// 			lastNumber:  "300",
// 			findNumber:  "70",
// 		},
// 		{
// 			startNumber: "1",
// 			lastNumber:  "300",
// 			findNumber:  "1",
// 		},
// 		{
// 			startNumber: "1",
// 			lastNumber:  "300",
// 			findNumber:  "299",
// 		},
// 		{
// 			startNumber: "1",
// 			lastNumber:  "300",
// 			findNumber:  "300",
// 		},
// 	}

// 	for _, tc := range tests {
// 		isNumberExist := suite.collector.IsTestApplicationNumberExist(tc.findNumber)
// 		start, last, _ := suite.collector.GetLastApplicationNumber(tc.startNumber, tc.lastNumber, isNumberExist)

// 		for {
// 			start, last, _ = suite.collector.GetLastApplicationNumber(start, last, isNumberExist)
// 			if last == start {
// 				break
// 			}
// 		}
// 	}
// }

// // Not used
// func (suite *CollectorTestSuite) TestFindNumberLogic() {
// 	type testcases struct {
// 		input  string
// 		result bool
// 	}
// 	isNumberExist := suite.collector.IsTestApplicationNumberExist("300")

// 	tests := []testcases{
// 		{
// 			input:  "301",
// 			result: false,
// 		},
// 		{
// 			input:  "300",
// 			result: true,
// 		},
// 		{
// 			input:  "299",
// 			result: true,
// 		},
// 		{
// 			input:  "1",
// 			result: true,
// 		},
// 		{
// 			input:  "500",
// 			result: false,
// 		},
// 	}

// 	for _, tc := range tests {
// 		isExist := isNumberExist(tc.input)
// 		suite.Equal(isExist, tc.result)
// 	}
// }

// // Not used
// func (suite *CollectorTestSuite) TestCollectorGetMidValue() {
// 	type testMidcases struct {
// 		start  int
// 		last   int
// 		result int
// 	}

// 	tests := []testMidcases{
// 		{
// 			start:  1,
// 			last:   9999999,
// 			result: 5000000,
// 		},
// 		{
// 			start:  1,
// 			last:   10,
// 			result: 5,
// 		},
// 		{
// 			start:  2,
// 			last:   10,
// 			result: 6,
// 		},
// 		{
// 			start:  3,
// 			last:   10,
// 			result: 6, // (10 - 3) / 2 + 3 => 6.5 => 버림 => 6
// 		},
// 		{
// 			start:  3,
// 			last:   11,
// 			result: 7,
// 		},
// 		{
// 			start:  250,
// 			last:   500,
// 			result: 375,
// 		},
// 		{
// 			start:  250,
// 			last:   499,
// 			result: 374,
// 		},
// 		{
// 			start:  150,
// 			last:   225,
// 			result: 187,
// 		},
// 	}

// 	for _, tc := range tests {
// 		mid := suite.collector.GetMidValue(tc.start, tc.last)
// 		suite.Equal(mid, tc.result)
// 	}
// }

// func (suite *CollectorTestSuite) TestCreateApplicationNumber() {
// 	type testMidcases struct {
// 		productCode  string
// 		year         string
// 		serialNumber int
// 		result       string
// 	}

// 	tests := []testMidcases{
// 		{
// 			productCode:  "40",
// 			year:         "2020",
// 			serialNumber: 1,
// 			result:       "4020200000001",
// 		},
// 		{
// 			productCode:  "40",
// 			year:         "1999",
// 			serialNumber: 9999999,
// 			result:       "4019999999999",
// 		},
// 		{
// 			productCode:  "40",
// 			year:         "1960",
// 			serialNumber: 1,
// 			result:       "4019600000001",
// 		},
// 	}

// 	for _, tc := range tests {
// 		mid := suite.collector.CreateApplicationNumber(tc.productCode, tc.year, tc.serialNumber)
// 		suite.Equal(mid, tc.result)
// 	}
// // }
// func (suite *CollectorTestSuite) TestAutoCrawler() {
// 	suite.collector.StartAutoCrawler()
// }

func TestCollectorSuite(t *testing.T) {
	suite.Run(t, new(CollectorTestSuite))
}

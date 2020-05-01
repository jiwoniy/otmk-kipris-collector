package collector

import (
	"fmt"
	"testing"

	"kipris-collector/model"
	"kipris-collector/types"

	"github.com/stretchr/testify/suite"
)

type CollectorTestSuite struct {
	suite.Suite
	collector types.Collector
}

func (suite *CollectorTestSuite) SetupTest() {
	config := collectorConfig{
		Endpoint:   "http://plus.kipris.or.kr/openapi/rest",
		AccessKey:  "=JbKg6deF5WolYTZcZkypzgLBbSVbjZC6VEgfccaQyw=",
		ListenAddr: ":8082",
	}

	collector, _ := NewCollector(config)
	suite.collector = collector
}

type testcases struct {
	url            string
	params         map[string]string
	dest           model.KiprisResponse
	responseStatus model.KiprisResponseStatus
}

func (suite *CollectorTestSuite) TestCollector() {
	parserInstance := suite.collector.GetParser()

	tests := []testcases{
		{
			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
			params: map[string]string{
				"applicationNumber": "4020200023099", // valid number
				"accessKey":         suite.collector.GetAccessKey(),
			},
			dest:           model.KiprisResponse{},
			responseStatus: model.Success,
		},
		{
			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
			params: map[string]string{
				"applicationNumber": "402020002309911", // invalid number
				"accessKey":         suite.collector.GetAccessKey(),
			},
			dest:           model.KiprisResponse{},
			responseStatus: model.Empty,
		},
		{
			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
			params: map[string]string{
				"applicationNumber": "", // invalid number
				"accessKey":         suite.collector.GetAccessKey(),
			},
			dest:           model.KiprisResponse{},
			responseStatus: model.Error,
		},
		{
			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
			params: map[string]string{
				"applicationNumber": "4020200023099", // valid number
				"accessKey":         "",
			},
			dest:           model.KiprisResponse{},
			responseStatus: model.Error,
		},

		{
			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
			params: map[string]string{
				"applicationNumber": "4020200023099", // valid number
				"accessKey":         suite.collector.GetAccessKey(),
			},
			dest:           model.KiprisResponse{},
			responseStatus: model.Success,
		},
		{
			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
			params: map[string]string{
				"applicationNumber": "402020002309911", // invalid number
				"accessKey":         suite.collector.GetAccessKey(),
			},
			dest:           model.KiprisResponse{},
			responseStatus: model.Empty,
		},
		{
			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
			params: map[string]string{
				"applicationNumber": "", // invalid number
				"accessKey":         suite.collector.GetAccessKey(),
			},
			dest:           model.KiprisResponse{},
			responseStatus: model.Error,
		},
		{
			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
			params: map[string]string{
				"applicationNumber": "4020200023099", // valid number
				"accessKey":         "",
			},
			dest:           model.KiprisResponse{},
			responseStatus: model.Error,
		},
	}

	for testIndex, tc := range tests {
		content, err := suite.collector.Get(tc.url, tc.params)

		// suite.Equal(nil, suite.VariableThatShouldStartAtFive)
		if err != nil {
			suite.Error(err)
		}

		parserInstance.Parse(content, &tc.dest)

		if tc.dest.Result() != tc.responseStatus {
			suite.Error(fmt.Errorf("This test index fail %d", testIndex))
		}
	}
}

func (suite *CollectorTestSuite) TestRealCollector() {
	params := map[string]string{
		"applicationNumber": "4020200000002", // valid number
		"accessKey":         suite.collector.GetAccessKey(),
	}

	content, err := suite.collector.Get("/trademarkInfoSearchService/applicationNumberSearchInfo", params)
	if err != nil {
		suite.Error(err)
	}

	var data model.KiprisResponse

	parseInstance := suite.collector.GetParser()
	parseInstance.Parse(content, &data)

	storage := suite.collector.GetStorage()

	storage.Create(&data.Body.Items.TradeMarkInfo)

	// ddd := data.(model.TradeMarkInfo)

	// tradeMarkInfo := model.TradeMarkInfo(data)
	// 	// storage.Create(&dest)
}

func TestCollectorSuite(t *testing.T) {
	suite.Run(t, new(CollectorTestSuite))
}

package collector

import (
	"testing"

	"kipris-collector/parser"
)

type testcases struct {
	url            string
	params         map[string]string
	dest           parser.Response
	responseStatus parser.KiprisResponseStatus
}

// func TestCollector(t *testing.T) {
// 	collector, _ := New()

// 	tests := []testcases{
// 		{
// 			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
// 			params: map[string]string{
// 				"applicationNumber": "4020200023099", // valid number
// 				"accessKey":         collector.GetAccessKey(),
// 			},
// 			dest:           parser.KiprisResponse{},
// 			responseStatus: parser.Success,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
// 			params: map[string]string{
// 				"applicationNumber": "402020002309911", // invalid number
// 				"accessKey":         collector.GetAccessKey(),
// 			},
// 			dest:           parser.KiprisResponse{},
// 			responseStatus: parser.Empty,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
// 			params: map[string]string{
// 				"applicationNumber": "", // invalid number
// 				"accessKey":         collector.GetAccessKey(),
// 			},
// 			dest:           parser.KiprisResponse{},
// 			responseStatus: parser.Error,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/applicationNumberSearchInfo",
// 			params: map[string]string{
// 				"applicationNumber": "4020200023099", // valid number
// 				"accessKey":         "",
// 			},
// 			dest:           parser.KiprisResponse{},
// 			responseStatus: parser.Error,
// 		},

// 		{
// 			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
// 			params: map[string]string{
// 				"applicationNumber": "4020200023099", // valid number
// 				"accessKey":         collector.GetAccessKey(),
// 			},
// 			dest:           parser.KiprisResponse{},
// 			responseStatus: parser.Success,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
// 			params: map[string]string{
// 				"applicationNumber": "402020002309911", // invalid number
// 				"accessKey":         collector.GetAccessKey(),
// 			},
// 			dest:           parser.KiprisResponse{},
// 			responseStatus: parser.Empty,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
// 			params: map[string]string{
// 				"applicationNumber": "", // invalid number
// 				"accessKey":         collector.GetAccessKey(),
// 			},
// 			dest:           parser.KiprisResponse{},
// 			responseStatus: parser.Error,
// 		},
// 		{
// 			url: "/trademarkInfoSearchService/trademarkDesignationGoodstInfo",
// 			params: map[string]string{
// 				"applicationNumber": "4020200023099", // valid number
// 				"accessKey":         "",
// 			},
// 			dest:           parser.KiprisResponse{},
// 			responseStatus: parser.Error,
// 		},
// 	}

// 	for testIndex, tc := range tests {
// 		err := collector.Get(tc.url, tc.params, &tc.dest)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		if tc.dest.Result() != tc.responseStatus {
// 			t.Errorf(fmt.Sprintf("This test index fail %d", testIndex))
// 		}
// 	}
// }

func TestRealCollector(t *testing.T) {
	collector, _ := New()

	params := map[string]string{
		"applicationNumber": "4020200000002", // valid number
		"accessKey":         collector.GetAccessKey(),
	}

	// dat := parserTypes.KiprisResponse{}

	content, err := collector.Get("/trademarkInfoSearchService/applicationNumberSearchInfo", params)
	if err != nil {
		t.Error(err)
	}

	var data parser.Response

	parseInstance := collector.GetParser()
	parseInstance.Parse(content, &data)

	storage := collector.GetStorage()

	storage.Create(&data.Body.Items.TradeMarkInfo)

	// ddd := data.(model.TradeMarkInfo)

	// tradeMarkInfo := model.TradeMarkInfo(data)
	// storage.Create(&dest)
}

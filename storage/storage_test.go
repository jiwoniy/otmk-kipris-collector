package storage

import (
	"testing"

	"github.com/jiwoniy/otmk-kipris-collector/model"
	"github.com/jiwoniy/otmk-kipris-collector/types"
)

// func TestKiprisApplicationNumber(t *testing.T) {
// 	storageConfig := types.StorageConfig{
// 		DbType:       "sqlite3",
// 		DbConnString: ":memory:",
// 	}

// 	storage, err := NewStorage(storageConfig)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	type testcases struct {
// 		data    model.KiprisApplicationNumber
// 		result  model.KiprisApplicationNumber
// 		success bool
// 	}

// 	tests := []testcases{
// 		{
// 			data: model.KiprisApplicationNumber{
// 				ApplicationNumber: "4020200000001",
// 				ProductCode:       "40",
// 				Year:              "2020",
// 				SerialNumber:      1,
// 			},
// 			result:  model.KiprisApplicationNumber{},
// 			success: true,
// 		},
// 		{
// 			data: model.KiprisApplicationNumber{
// 				ApplicationNumber: "4020000000001",
// 				ProductCode:       "40",
// 				Year:              "2000",
// 				SerialNumber:      1,
// 			},
// 			result:  model.KiprisApplicationNumber{},
// 			success: true,
// 		},
// 		{
// 			data: model.KiprisApplicationNumber{
// 				ApplicationNumber: "4120100000002",
// 				ProductCode:       "41",
// 				Year:              "2010",
// 				SerialNumber:      2,
// 			},
// 			result:  model.KiprisApplicationNumber{},
// 			success: true,
// 		},
// 		{
// 			data: model.KiprisApplicationNumber{
// 				ApplicationNumber: "4519990999990",
// 				ProductCode:       "45",
// 				Year:              "1999",
// 				SerialNumber:      999990,
// 			},
// 			result:  model.KiprisApplicationNumber{},
// 			success: true,
// 		},
// 		{
// 			data: model.KiprisApplicationNumber{
// 				ApplicationNumber: "4020000000001",
// 				ProductCode:       "40",
// 				Year:              "2000",
// 				SerialNumber:      1,
// 			},
// 			result:  model.KiprisApplicationNumber{},
// 			success: true,
// 		},
// 		{
// 			data: model.KiprisApplicationNumber{
// 				ApplicationNumber: "4020209999999",
// 				ProductCode:       "40",
// 				Year:              "2020",
// 				SerialNumber:      9999999,
// 			},
// 			result:  model.KiprisApplicationNumber{},
// 			success: false,
// 		},
// 	}

// 	for tcIndex, tc := range tests {
// 		if tc.success {
// 			storage.Create(&tc.data)
// 		}

// 		storage.GetKiprisApplicationNumber(tc.data, &tc.result)
// 		if tc.success == true && tc.result.ApplicationNumber == "" {
// 			t.Errorf("testcase %d error: %s", tcIndex+1, err)
// 		} else if tc.success == false && tc.result.ApplicationNumber != "" {
// 			t.Errorf("testcase %d error: %s", tcIndex+1, err)
// 		}
// 	}

// 	searchResult := make([]model.KiprisApplicationNumber, 0)
// 	searchData := model.KiprisApplicationNumber{
// 		ProductCode: "40",
// 	}
// 	storage.GetKiprisApplicationNumberList(searchData, &searchResult, 1, 20)

// 	if len(searchResult) != 3 {
// 		t.Errorf("search data find %v shoud be %d", searchData, len(searchResult))
// 	}

// 	searchData = model.KiprisApplicationNumber{
// 		Year: "2000",
// 	}
// 	storage.GetKiprisApplicationNumberList(searchData, &searchResult, 1, 20)
// 	if len(searchResult) != 2 {
// 		t.Errorf("search data find %v shoud be %d", searchData, len(searchResult))
// 	}

// }

func TestTradeMarkInfo(t *testing.T) {
	storageConfig := types.StorageConfig{
		DbType:       "sqlite3",
		DbConnString: ":memory:",
	}
	storage, err := NewStorage(storageConfig)
	if err != nil {
		t.Error(err)
	}

	type testcases struct {
		data    model.TradeMarkInfo
		result  model.TradeMarkInfo
		success bool
	}

	tests := []testcases{
		{
			data: model.TradeMarkInfo{
				SerialNumber:       "1",
				ApplicationNumber:  "4020190000001",
				AppReferenceNumber: "",
				ApplicationDate:    "20190101",
				PublicNumber:       "4020190084686",

				PublicDate:               "20190731",
				RegistrationPublicNumber: "4020190125782",
				RegistrationPublicDate:   "20191230",
				RegistrationNumber:       "4015569910000",
				RegReferenceNumber:       "",

				RegistrationDate:       "20191223",
				PriorityClaimNumber:    "",
				PriorityClaimDate:      "",
				ApplicationStatus:      "등록",
				GoodClassificationCode: "43",

				ViennaCode:                  "",
				ApplicantName:               "김현호",
				AgentName:                   "",
				RegistrationRightholderName: "김현호 kimHyunHo",
				Title:                       "돈퍼",

				FulltextExistFlag: "Y",
				ImagePath:         "http://plus.kipris.or.kr/kiprisplusws/fileToss.jsp?arg=ed43a0609e94d6e251697a9d72a9134435594e3384b42e76c0c1194596d9ce713eab8b480c0bac7f959eac6a27cc4b4806ef5b7e099c903d",
				ThumbnailPath:     "http://plus.kipris.or.kr/kiprisplusws/fileToss.jsp?arg=ad7a17eeeef6e4ea4b5e22ef00dd3e293e70a322c3ead7b643de6089e8a3e092c32a212d34313535ec88eb55c218f60067d17b9d393c7bf1",
			},
			success: true,
		},
		{
			// duplicate error
			data: model.TradeMarkInfo{
				SerialNumber:       "1",
				ApplicationNumber:  "4020190000001",
				AppReferenceNumber: "",
				ApplicationDate:    "20190101",
				PublicNumber:       "4020190084686",

				PublicDate:               "20190731",
				RegistrationPublicNumber: "4020190125782",
				RegistrationPublicDate:   "20191230",
				RegistrationNumber:       "4015569910000",
				RegReferenceNumber:       "",

				RegistrationDate:       "20191223",
				PriorityClaimNumber:    "",
				PriorityClaimDate:      "",
				ApplicationStatus:      "등록",
				GoodClassificationCode: "43",

				ViennaCode:                  "",
				ApplicantName:               "김현호",
				AgentName:                   "",
				RegistrationRightholderName: "김현호 kimHyunHo",
				Title:                       "돈퍼",

				FulltextExistFlag: "Y",
				ImagePath:         "http://plus.kipris.or.kr/kiprisplusws/fileToss.jsp?arg=ed43a0609e94d6e251697a9d72a9134435594e3384b42e76c0c1194596d9ce713eab8b480c0bac7f959eac6a27cc4b4806ef5b7e099c903d",
				ThumbnailPath:     "http://plus.kipris.or.kr/kiprisplusws/fileToss.jsp?arg=ad7a17eeeef6e4ea4b5e22ef00dd3e293e70a322c3ead7b643de6089e8a3e092c32a212d34313535ec88eb55c218f60067d17b9d393c7bf1",
			},
			success: false,
		},
	}

	for tcIndex, tc := range tests {
		err = storage.Create(&tc.data)

		if tc.success == true && err != nil {
			t.Errorf("testcase %d error: %s", tcIndex+1, err)
		} else if tc.success == false && err == nil {
			t.Errorf("testcase %d error", tcIndex+1)
		} else {
			storage.GetTradeMarkInfo(model.TradeMarkInfo{
				SerialNumber:      tc.data.SerialNumber,
				ApplicationNumber: tc.data.ApplicationNumber,
			}, &tc.result)
			if tc.result.ApplicationNumber != tc.data.ApplicationNumber {
				t.Errorf("testcase %d error", tcIndex+1)
			}
		}
	}
}

func TestTrademarkDesignationGoodstInfo(t *testing.T) {
	storageConfig := types.StorageConfig{
		DbType:       "sqlite3",
		DbConnString: ":memory:",
	}

	storage, err := NewStorage(storageConfig)
	if err != nil {
		t.Error(err)
	}

	type testcases struct {
		data    model.TrademarkDesignationGoodstInfo
		result  []model.TrademarkDesignationGoodstInfo
		success bool
	}

	trademarkInfo := model.TradeMarkInfo{
		SerialNumber:      "1",
		ApplicationNumber: "4020190000001",
	}

	tests := []testcases{
		{
			data: model.TrademarkDesignationGoodstInfo{
				ApplicationNumber:                             trademarkInfo.ApplicationNumber,
				DesignationGoodsSerialNumber:                  "1",
				DesignationGoodsClassificationInformationCode: "43",
				SimilargroupCode:                              "G0301",
				DesignationGoodsHangeulName:                   "중국식퓨전음식전문점체인업",
				DesignationGoodsEnglishsentenceName:           "",
			},
			success: true,
		},
		{
			data: model.TrademarkDesignationGoodstInfo{
				ApplicationNumber:                             trademarkInfo.ApplicationNumber,
				DesignationGoodsSerialNumber:                  "2",
				DesignationGoodsClassificationInformationCode: "43",
				SimilargroupCode:                              "G0502",
				DesignationGoodsHangeulName:                   "중국식퓨전음식전문점체인업",
				DesignationGoodsEnglishsentenceName:           "",
			},
			success: true,
		},
		{
			data: model.TrademarkDesignationGoodstInfo{
				ApplicationNumber:                             trademarkInfo.ApplicationNumber,
				DesignationGoodsSerialNumber:                  "3",
				DesignationGoodsClassificationInformationCode: "43",
				SimilargroupCode:                              "S120602",
				DesignationGoodsHangeulName:                   "중국식퓨전음식전문점체인업",
				DesignationGoodsEnglishsentenceName:           "",
			},
			success: true,
		},
		{
			data: model.TrademarkDesignationGoodstInfo{
				ApplicationNumber:                             "",
				DesignationGoodsSerialNumber:                  "3",
				DesignationGoodsClassificationInformationCode: "43",
				SimilargroupCode:                              "S120602",
				DesignationGoodsHangeulName:                   "중국식퓨전음식전문점체인업",
				DesignationGoodsEnglishsentenceName:           "",
			},
			success: false,
		},
		{
			data: model.TrademarkDesignationGoodstInfo{
				ApplicationNumber:                             trademarkInfo.ApplicationNumber,
				DesignationGoodsSerialNumber:                  "",
				DesignationGoodsClassificationInformationCode: "43",
				SimilargroupCode:                              "S120602",
				DesignationGoodsHangeulName:                   "중국식퓨전음식전문점체인업",
				DesignationGoodsEnglishsentenceName:           "",
			},
			success: false,
		},
		{
			data: model.TrademarkDesignationGoodstInfo{
				ApplicationNumber:                             trademarkInfo.ApplicationNumber,
				DesignationGoodsSerialNumber:                  "3",
				DesignationGoodsClassificationInformationCode: "",
				SimilargroupCode:                              "S120602",
				DesignationGoodsHangeulName:                   "중국식퓨전음식전문점체인업",
				DesignationGoodsEnglishsentenceName:           "",
			},
			success: false,
		},
		{
			data: model.TrademarkDesignationGoodstInfo{
				ApplicationNumber:                             trademarkInfo.ApplicationNumber,
				DesignationGoodsSerialNumber:                  "dkd",
				DesignationGoodsClassificationInformationCode: "43",
				SimilargroupCode:                              "",
				DesignationGoodsHangeulName:                   "중국식퓨전음식전문점체인업",
				DesignationGoodsEnglishsentenceName:           "",
			},
			success: false,
		},
	}

	for tcIndex, tc := range tests {
		err = storage.Create(&tc.data)
		if tc.success == true && err != nil {
			t.Errorf("testcase %d error: %s", tcIndex+1, err)
		} else if tc.success == false && err == nil {
			t.Errorf("testcase %d error: %s", tcIndex+1, err)
		}
		if tc.success == true && err == nil {
			storage.GetTrademarkDesignationGoodstInfo(model.TrademarkDesignationGoodstInfo{
				ApplicationNumber:                             tc.data.ApplicationNumber,
				DesignationGoodsSerialNumber:                  tc.data.DesignationGoodsSerialNumber,
				DesignationGoodsClassificationInformationCode: tc.data.DesignationGoodsClassificationInformationCode,
				SimilargroupCode:                              tc.data.SimilargroupCode,
				DesignationGoodsHangeulName:                   tc.data.DesignationGoodsHangeulName,
			}, &tc.result)
			if tc.result[0].ApplicationNumber != tc.data.ApplicationNumber ||
				tc.result[0].DesignationGoodsSerialNumber != tc.data.DesignationGoodsSerialNumber ||
				tc.result[0].DesignationGoodsClassificationInformationCode != tc.data.DesignationGoodsClassificationInformationCode ||
				tc.result[0].DesignationGoodsHangeulName != tc.data.DesignationGoodsHangeulName ||
				tc.result[0].SimilargroupCode != tc.data.SimilargroupCode {
				t.Errorf("testcase %d error", tcIndex+1)
			}
		}
	}
}

func TestKiprisCollector(t *testing.T) {
	storageConfig := types.StorageConfig{
		DbType:       "sqlite3",
		DbConnString: ":memory:",
	}

	storage, err := NewStorage(storageConfig)
	if err != nil {
		t.Error(err)
	}

	type testcases struct {
		data    model.KiprisCollectorStatus
		result  model.KiprisCollectorStatus
		success bool
	}

	tests := []testcases{
		{
			data: model.KiprisCollectorStatus{
				ApplicationNumber:                  "402020000001",
				TradeMarkInfoStatus:                model.Success,
				TradeMarkDesignationGoodInfoStatus: model.Success,
			},
			result:  model.KiprisCollectorStatus{},
			success: true,
		},
		{
			data: model.KiprisCollectorStatus{
				ApplicationNumber:                  "402020000002",
				TradeMarkInfoStatus:                model.Error,
				TradeMarkDesignationGoodInfoStatus: model.Error,
			},
			result:  model.KiprisCollectorStatus{},
			success: true,
		},
		{
			data: model.KiprisCollectorStatus{
				ApplicationNumber:                  "402020000003",
				TradeMarkInfoStatus:                model.Empty,
				TradeMarkDesignationGoodInfoStatus: model.Empty,
			},
			result:  model.KiprisCollectorStatus{},
			success: true,
		},
	}

	for tcIndex, tc := range tests {
		err = storage.Create(&tc.data)
		// storage.GetKiprisApplicationNumber(tc.data, &tc.result)
		if tc.success == true && err != nil {
			t.Errorf("testcase %d error: %s", tcIndex+1, err)
		}
	}
}

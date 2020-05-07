package parser

import (
	"reflect"
	"testing"

	"github.com/jiwoniy/otmk-kipris-collector/model"
)

type testcases struct {
	input          string
	success        bool
	data           interface{}
	responseStatus model.KiprisResponseStatus
}

func TestParserInstance(t *testing.T) {
	tests := []testcases{
		{
			input:   "xml",
			success: true,
		},
		{
			input:   "json",
			success: false,
		},
		{
			input:   "test",
			success: false,
		},
	}

	for _, tc := range tests {
		_, err := NewParser(tc.input)
		if tc.success == true && err != nil {
			t.Error(err)
		} else if tc.success == false && err == nil {
			t.Error(err)
		}
	}
}

func TestRead(t *testing.T) {
	parser, _ := NewParser("xml")
	tests := []testcases{
		{
			input:   "./test_data/TrademarkDesignationGoodstInfo.xml",
			success: true,
		},
		{
			input:   "./test_data/notExistFile.xml",
			success: false,
		},
		{
			input:   "./test_data/TradeMarkInfo.xml",
			success: true,
		},
	}

	for _, tc := range tests {
		_, err := parser.Read(tc.input)
		if tc.success == true && err != nil {
			t.Error(err)
		}
		//fmt.Printf("File contents: %s", content)
	}
}

func TestParsing(t *testing.T) {
	parser, _ := NewParser("xml")
	tests := []testcases{
		{
			input:          "./test_data/TrademarkDesignationGoodstInfo.xml",
			responseStatus: model.Success,
			// success: true,
			data: model.KiprisResponse{
				Header: model.Header{
					ResultCode: "",
					ResultMsg:  "",
				},
				Body: model.Body{
					Items: model.Items{
						TrademarkDesignationGoodstInfo: []model.TrademarkDesignationGoodstInfo{
							{
								DesignationGoodsSerialNumber:                  "1",
								DesignationGoodsClassificationInformationCode: "43",
								SimilargroupCode:                              "G0301",
								DesignationGoodsHangeulName:                   "중국식퓨전음식전문점체인업",
							},
							{
								DesignationGoodsSerialNumber:                  "2",
								DesignationGoodsClassificationInformationCode: "43",
								SimilargroupCode:                              "G0502",
								DesignationGoodsHangeulName:                   "중국식퓨전음식전문점체인업",
							},
							{
								DesignationGoodsSerialNumber:                  "3",
								DesignationGoodsClassificationInformationCode: "43",
								SimilargroupCode:                              "S120602",
								DesignationGoodsHangeulName:                   "중국식퓨전음식전문점체인업",
							},
							{
								DesignationGoodsSerialNumber:                  "4",
								DesignationGoodsClassificationInformationCode: "43",
								SimilargroupCode:                              "G0301",
								DesignationGoodsHangeulName:                   "중국음식점업",
							},
							{
								DesignationGoodsSerialNumber:                  "5",
								DesignationGoodsClassificationInformationCode: "43",
								SimilargroupCode:                              "G0502",
								DesignationGoodsHangeulName:                   "중국음식점업",
							},
							{
								DesignationGoodsSerialNumber:                  "6",
								DesignationGoodsClassificationInformationCode: "43",
								SimilargroupCode:                              "S120602",
								DesignationGoodsHangeulName:                   "중국음식점업",
							},
							{
								DesignationGoodsSerialNumber:                  "7",
								DesignationGoodsClassificationInformationCode: "43",
								SimilargroupCode:                              "G0301",
								DesignationGoodsHangeulName:                   "중국식퓨전음식전문점경영업",
							},
							{
								DesignationGoodsSerialNumber:                  "8",
								DesignationGoodsClassificationInformationCode: "43",
								SimilargroupCode:                              "G0502",
								DesignationGoodsHangeulName:                   "중국식퓨전음식전문점경영업",
							},
							{
								DesignationGoodsSerialNumber:                  "9",
								DesignationGoodsClassificationInformationCode: "43",
								SimilargroupCode:                              "S120602",
								DesignationGoodsHangeulName:                   "중국식퓨전음식전문점경영업",
							},
						},
					},
				},
			},
		},
		{
			input:          "./test_data/TradeMarkInfo.xml",
			responseStatus: model.Success,
			data: model.KiprisResponse{
				Header: model.Header{
					ResultCode: "",
					ResultMsg:  "",
				},
				Body: model.Body{
					Items: model.Items{
						TotalSearchCount: "1",
						TradeMarkInfo: model.TradeMarkInfo{
							SerialNumber:       "1",
							ApplicationNumber:  "4020190000011",
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
					},
				},
			},
		},
		{
			input:          "./test_data/Empty_TradeMarkInfo.xml",
			responseStatus: model.Empty,
			data: model.KiprisResponse{
				Header: model.Header{
					ResultCode: "",
					ResultMsg:  "",
				},
				Body: model.Body{
					Items: model.Items{
						TotalSearchCount: "0",
					},
				},
			},
		},
		{
			input:          "./test_data/Empty_TrademarkDesignationGoodstInfo.xml",
			responseStatus: model.Empty,
			data: model.KiprisResponse{
				Header: model.Header{
					ResultCode: "",
					ResultMsg:  "",
				},
				Body: model.Body{
					Items: model.Items{},
				},
			},
		},
		{
			input:          "./test_data/Error_InvalidKey.xml",
			responseStatus: model.Error,
			data: model.KiprisResponse{
				Header: model.Header{
					ResultCode: "30",
					// ResultMsg:  "등록된 서비스키를 입력해 주십시오(Access key &amp; Service key is not registerd error)",
					ResultMsg: "",
				},
				// Body: Body{
				// 	Items: Items{
				// 		TotalSearchCount: "",
				// 	},
				// },
			},
		},
		{
			input:          "./test_data/Error_EmptyApplicationNumber.xml",
			responseStatus: model.Error,
			data: model.KiprisResponse{
				Header: model.Header{
					ResultCode: "10",
					ResultMsg:  "",
				},
				// Body: Body{
				// 	Items: Items{
				// 		TotalSearchCount: "",
				// 	},w
				// },
			},
		},
	}

	for _, tc := range tests {
		content, err := parser.Read(tc.input)
		if tc.success == true && err != nil {
			t.Error(err)
		}
		var data model.KiprisResponse
		parser.Parse(content, &data)

		isSame := reflect.DeepEqual(tc.data, data)
		if !isSame || data.Result() != tc.responseStatus {
			t.Errorf("file %s parsing fail", tc.input)
		}
	}
}

package storage

import (
	"kipris-collector/model"
	"testing"
)

func TestTradeMarkInfoCreate(t *testing.T) {
	type testcases struct {
		data    model.TradeMarkInfo
		success bool
	}

	storageConfig := StorageConfig{
		DbType:       "sqlite3",
		DbConnString: ":memory:",
	}
	storage, err := NewStorage(storageConfig)
	if err != nil {
		t.Error(err)
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
		}

		if tc.success == false && err == nil {
			t.Errorf("testcase %d error", tcIndex+1)
		}
	}

}

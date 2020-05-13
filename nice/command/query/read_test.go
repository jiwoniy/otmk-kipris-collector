package query

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestSearchResult(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./../nice_code.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	client := NewKeeper(db)

	type testCase struct {
		Text               string
		ClassificationCode string
		TotalRecord        int
	}
	var tcase []testCase
	tcase = append(tcase, testCase{
		Text:               "",
		ClassificationCode: "01",
		TotalRecord:        26,
	}, testCase{
		Text:               "",
		ClassificationCode: "02",
		TotalRecord:        8,
	},
		testCase{
			Text:               "휴대폰",
			ClassificationCode: "02",
			TotalRecord:        0,
		},
		testCase{
			Text:               "휴대폰",
			ClassificationCode: "",
			TotalRecord:        15,
		},
		testCase{
			Text:               "휴대폰",
			ClassificationCode: "01",
			TotalRecord:        1,
		})

	for _, tc := range tcase {
		result, _ := client.SearchSimilarGroups(tc.Text, tc.ClassificationCode, 2, 20)
		if result.TotalRecord != tc.TotalRecord {
			t.Errorf("Search %s with %s; want: %d but result: %d", tc.Text, tc.ClassificationCode, tc.TotalRecord, result.TotalRecord)
		}
	}

	if err != nil {
		panic(err)
	}
}

func TestSearchText(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./../nice_code.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	client := NewKeeper(db)

	type testCase struct {
		Id     string
		Result string
	}
	var tcase []testCase
	tcase = append(tcase, testCase{
		Id:     "74",
		Result: "공업용 화학제 및 접착제",
	})

	for _, tc := range tcase {
		result, _ := client.GetSimilarCodeText(tc.Id)
		if result != tc.Result {
			t.Errorf("Get text; want: %s but result: %s", tc.Result, result)
		}
	}

	if err != nil {
		panic(err)
	}
}

func TestSearchPaging(t *testing.T) {
	db, err := gorm.Open("sqlite3", "./../nice_code.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	client := NewKeeper(db)

	type testCase struct {
		Text               string
		ClassificationCode string
		Page               int
		Size               int
		TotalRecord        int
		TotalPage          int
	}
	var tcase []testCase
	tcase = append(tcase,
		testCase{
			Text:               "",
			ClassificationCode: "01",
			Page:               1,
			Size:               26,
			TotalRecord:        26,
			TotalPage:          6,
		},
		testCase{
			Text:               "",
			ClassificationCode: "01",
			Page:               1,
			Size:               5,
			TotalRecord:        26,
			TotalPage:          6,
		},
		testCase{
			Text:               "",
			ClassificationCode: "01",
			Page:               2,
			Size:               5,
			TotalRecord:        26,
			TotalPage:          6,
		},
		testCase{
			Text:               "",
			ClassificationCode: "01",
			Page:               3,
			Size:               5,
			TotalRecord:        26,
			TotalPage:          6,
		},
		testCase{
			Text:               "",
			ClassificationCode: "01",
			Page:               4,
			Size:               5,
			TotalRecord:        26,
			TotalPage:          6,
		},
		testCase{
			Text:               "",
			ClassificationCode: "01",
			Page:               5,
			Size:               5,
			TotalRecord:        26,
			TotalPage:          6,
		},
		testCase{
			Text:               "",
			ClassificationCode: "01",
			Page:               6,
			Size:               5,
			TotalRecord:        26,
			TotalPage:          6,
		},
		testCase{
			Text:               "",
			ClassificationCode: "01",
			Page:               7,
			Size:               5,
			TotalRecord:        26,
			TotalPage:          6,
		},
	)

	for _, tc := range tcase {
		result, _ := client.SearchSimilarGroups(tc.Text, tc.ClassificationCode, tc.Size, tc.Page)
		if result.TotalPage == tc.TotalPage && result.Page != tc.Page {
			t.Errorf("Search %s with %s; want: %d but result: %d", tc.Text, tc.ClassificationCode, tc.TotalRecord, result.TotalRecord)
		}
	}

	if err != nil {
		panic(err)
	}
}

package storage

import (
	"kipris-collector/model"
	"testing"
	"time"
)

// type testcases struct {
// 	url            string
// 	params         map[string]string
// 	data           parser.KiprisResponse
// 	responseStatus parser.KiprisResponseStatus
// }

// func TestStorage(t *testing.T) {
// }

func TestCreate(t *testing.T) {
	storage, err := New()
	if err != nil {
		t.Error(err)
	}

	tradeMarkInfo := model.TradeMarkInfo{
		SerialNumber:       "1",
		ApplicationNumber:  "4020190000011b",
		AppReferenceNumber: time.Now().String(),
		ApplicationDate:    "20190101",
	}

	err = storage.Create(&tradeMarkInfo)
	if err != nil {
		t.Error(err)
	}

}

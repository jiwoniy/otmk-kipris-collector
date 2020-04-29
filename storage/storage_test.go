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

// TODO
func TestCreate(t *testing.T) {
	storage, err := New()
	if err != nil {
		t.Error(err)
	}

	tradeMarkInfo := model.TradeMarkInfo{
		SerialNumber:       "1",
		ApplicationNumber:  "1112211",
		AppReferenceNumber: time.Now().String(),
		ApplicationDate:    "20190101",
	}

	err = storage.Create(&tradeMarkInfo)
	if err != nil {
		t.Error(err)
	}
}

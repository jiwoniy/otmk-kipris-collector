package model

import (
	"encoding/xml"

	// "fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

type KiprisResponseStatus int

const (
	Error KiprisResponseStatus = iota + 1 // 1
	Success
	Empty
)

type TrimString string

var validate *validator.Validate

// custom xml string for whitespace trim
func (str *TrimString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	trimStr := strings.TrimSpace(s)
	*str = TrimString(trimStr)

	return nil
}

func (str TrimString) Valid() bool {
	if len(str) < 1 {
		return false
	}
	return true
}

type KiprisResponse struct {
	// XMLName xml.Name `xml:"response"`
	Header Header `xml:"header"`
	Body   Body   `xml:"body"`
}

func (res *KiprisResponse) Result() KiprisResponseStatus {
	result := *res
	resultCode := result.Header.ResultCode

	if resultCode != "" {
		return Error
	} else if reflect.DeepEqual(KiprisResponse{
		Header: Header{
			ResultCode: "",
			ResultMsg:  "",
		},
		Body: Body{
			Items: Items{
				TotalSearchCount: "0",
			},
		},
	}, *res) {
		return Empty
	} else if reflect.DeepEqual(KiprisResponse{
		Header: Header{
			ResultCode: "",
			ResultMsg:  "",
		},
		Body: Body{
			Items: Items{},
		},
	}, *res) {
		return Empty
	}

	return Success
}

type Items struct {
	// XMLName                        xml.Name                         `xml:"items"`
	TrademarkDesignationGoodstInfo []TrademarkDesignationGoodstInfo `xml:"trademarkDesignationGoodstInfo"`
	TradeMarkInfo                  TradeMarkInfo                    `xml:"TradeMarkInfo"`
	TotalSearchCount               TotalSearchCount                 `xml:"TotalSearchCount"`
}

type Body struct {
	// XMLName xml.Name `xml:"body"`
	Items Items `xml:"items"`
}

type Header struct {
	// XMLName    xml.Name `xml:"header"`
	ResultCode TrimString `xml:"resultCode"`
	ResultMsg  TrimString `xml:"resultMsg"`
}

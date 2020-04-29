package parser

import (
	"encoding/xml"
	"reflect"
	"strings"
)

type KiprisResponseStatus int

const (
	Error KiprisResponseStatus = iota
	Success
	Empty
)

type TrimString string

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

	// fmt.Println(data)
	return Success
}

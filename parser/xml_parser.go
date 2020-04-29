package parser

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
)

type xmlParser struct{}

func (p *xmlParser) Read(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Please use valid file name %s is not exist", filename))
	}

	return content, nil
}

func (p *xmlParser) Parse(data []byte, v interface{}) error {
	err := xml.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

func (p *xmlParser) Print(v interface{}) {
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}

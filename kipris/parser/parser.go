package parser

import (
	"errors"
	"fmt"

	"github.com/jiwoniy/otmk-kipris-collector/kipris/types"
)

func NewParser(parserType string) (types.Parser, error) {
	switch parserType {
	case "xml":
		{
			return &xmlParser{}, nil
		}
	default:
		{
			return nil, errors.New(fmt.Sprintf("%s type is not support", parserType))
		}
	}
}

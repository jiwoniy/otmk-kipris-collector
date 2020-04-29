package parser

import (
	"errors"
	"fmt"

	"kipris-collector/types"
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

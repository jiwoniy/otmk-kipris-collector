package collector

import (
	"kipris-collector/parser"
	"kipris-collector/types"
	"kipris-collector/utils"
)

type kiprisCollector struct {
	endpt     string
	accessKey string
	parser    types.Parser
}

func NewCollector(config collectorConfig) (types.Collector, error) {
	parserInstance, err := parser.NewParser("xml")
	if err != nil {
		return nil, err
	}
	return &kiprisCollector{
		endpt:     config.Endpoint,
		accessKey: config.AccessKey,
		parser:    parserInstance,
	}, nil
}

func (c *kiprisCollector) GetEndpoint() string {
	return c.endpt
}

func (c *kiprisCollector) GetAccessKey() string {
	return c.accessKey
}

func (c *kiprisCollector) GetParser() types.Parser {
	return c.parser
}

func (c *kiprisCollector) Get(url string, params map[string]string, dest interface{}) error {
	caller, err := utils.BuildRESTCaller(c.endpt).Build()
	if err != nil {
		return err
	}

	err = caller.Get(url, params, nil, &dest)

	if err != nil {
		return err
	}

	return err
}

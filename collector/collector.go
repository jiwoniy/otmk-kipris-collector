package collector

import (
	"kipris-collector/parse"
	"kipris-collector/utils"
)

type kiprisCollector struct {
	endpt     string
	accessKey string
	parser    parse.Parse
}

func New(config collectorConfig) *kiprisCollector {
	return &kiprisCollector{
		endpt:     config.Endpoint,
		accessKey: config.AccessKey,
		parser:    parse.NewParse("xml"),
	}
}

func (c *kiprisCollector) Get(url string) error {
	caller, err := utils.BuildRESTCaller(c.endpt).Build()
	if err != nil {
		return err
	}

	param := map[string]string{
		"applicationNumber": "4020200023099",
		"accessKey":         c.accessKey,
	}

	var data parse.Response

	err = caller.Get(url, param, nil, &data)

	if err != nil {
		return err
	}

	c.parser.Print(data)

	return err
}

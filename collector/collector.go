package collector

import (
	"fmt"
	"kipris-collector/model"
	"kipris-collector/parser"
	"kipris-collector/storage"
	"kipris-collector/types"
	"kipris-collector/utils"
)

type kiprisCollector struct {
	endpt     string
	accessKey string
	parser    types.Parser
	storage   types.Storage
}

func NewCollector(config collectorConfig) (types.Collector, error) {
	parserInstance, err := parser.NewParser("xml")

	if err != nil {
		return nil, err
	}

	storage, err := storage.NewStorage(storage.StorageConfig{
		DbType:       config.DbType,
		DbConnString: config.DbConnString,
	})

	if err != nil {
		return nil, err
	}

	return &kiprisCollector{
		endpt:     config.Endpoint,
		accessKey: config.AccessKey,
		parser:    parserInstance,
		storage:   storage,
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

func (c *kiprisCollector) GetStorage() types.Storage {
	return c.storage
}

func (c *kiprisCollector) Get(url string, params map[string]string) ([]byte, error) {
	caller, err := utils.BuildRESTCaller(c.endpt).Build()
	if err != nil {
		return nil, err
	}

	body, err := caller.Get(nil, url, params)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *kiprisCollector) GetApplicationNumber(applicationNumber string) bool {
	params := map[string]string{
		"applicationNumber": applicationNumber,
		"accessKey":         c.GetAccessKey(),
	}
	content, err := c.Get("/trademarkInfoSearchService/applicationNumberSearchInfo", params)

	if err != nil {
		fmt.Println(err)
		return false
	}

	var tradeMarkInfo model.KiprisResponse
	c.parser.Parse(content, &tradeMarkInfo)
	// fmt.Println(tradeMarkInfo.Result())
	c.storage.Create(&tradeMarkInfo.Body.Items.TradeMarkInfo)

	content, err = c.Get("/trademarkInfoSearchService/trademarkDesignationGoodstInfo", params)
	if err != nil {
		fmt.Println(err)
		return false
	}

	var trademarkDesignationGoodstInfo model.KiprisResponse
	c.parser.Parse(content, &trademarkDesignationGoodstInfo)

	// fmt.Println(trademarkDesignationGoodstInfo.Result())

	for _, good := range trademarkDesignationGoodstInfo.Body.Items.TrademarkDesignationGoodstInfo {
		good.ApplicationNumber = applicationNumber
		err := c.storage.Create(&good)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}

	statistic := model.KiprisCollector{
		ApplicationNumber: applicationNumber,
		Status:            1,
		// Status:            tradeMarkInfo.Result() && trademarkDesignationGoodstInfo.Result(),
	}
	c.storage.Create(&statistic)
	return true
}

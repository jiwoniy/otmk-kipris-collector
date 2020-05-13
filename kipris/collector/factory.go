package collector

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"github.com/jiwoniy/otmk-kipris-collector/kipris/types"
)

func New() (types.Collector, error) {
	configPath := flag.String("cfg", "./config.json", "path to the configuration file")
	flag.Parse()

	var cfg types.CollectorConfig

	cfgData, err := ioutil.ReadFile(*configPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(cfgData, &cfg); err != nil {
		return nil, err
	}

	c, err := NewCollector(cfg)
	if err != nil {
		return nil, err
	}
	return c, nil
}

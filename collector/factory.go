package collector

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"github.com/jiwoniy/otmk-kipris-collector/types"
)

type CollectorConfig struct {
	Endpoint     string `json:"endpoint"`
	AccessKey    string `json:"access_key"`
	ListenAddr   string `json:"listen_addr"`
	DbType       string `json:"dbType"`
	DbConnString string `json:"DbConnString"`
}

func New() (types.Collector, error) {
	configPath := flag.String("cfg", "./config.json", "path to the configuration file")
	flag.Parse()

	var cfg CollectorConfig

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

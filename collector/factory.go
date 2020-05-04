package collector

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"kipris-collector/types"
)

type collectorConfig struct {
	Endpoint     string `json:"endpoint"`
	AccessKey    string `json:"access_key"`
	ListenAddr   string `json:"listen_addr"`
	DbType       string `json:"dbType"`
	DbConnString string `json:"DbConnString"`
}

func New() (types.Collector, error) {
	configPath := flag.String("cfg", "./config.json", "path to the configuration file")
	flag.Parse()

	var cfg collectorConfig

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

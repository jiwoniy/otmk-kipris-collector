package storage

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"kipris-collector/types"
)

type StorageConfig struct {
	// base on gorm format
	DbType       string `json:"dbType"`
	DbConnString string `json:"dbConnString"`
}

func New() (types.Storage, error) {
	configPath := flag.String("cfg", "./config.json", "path to the configuration file")
	flag.Parse()

	var cfg StorageConfig

	cfgData, err := ioutil.ReadFile(*configPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(cfgData, &cfg); err != nil {
		return nil, err
	}

	if c, err := NewStorage(cfg); err != nil {
		return nil, err
	} else {
		return c, nil
	}
}

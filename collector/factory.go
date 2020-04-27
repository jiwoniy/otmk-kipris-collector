package collector

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

type collectorConfig struct {
	Endpoint   string `json:"endpoint"`
	AccessKey  string `json:"access_key"`
	ListenAddr string `json:"listen_addr"`
}

func NewCollector() (*kiprisCollector, error) {
	configPath := flag.String("cfg", "./config.json", "path to the configuration file")
	flag.Parse()

	var cfg collectorConfig

	cfgData, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(cfgData, &cfg); err != nil {
		panic(err)
	}

	c := New(cfg)
	return c, nil
}

// func main() {
// 	configPath := flag.String("cfg", "./config.json", "path to the configuration file")
// 	flag.Parse()

// 	var cfg sdk.AppConfig

// 	cfgData, err := ioutil.ReadFile(*configPath)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := json.Unmarshal(cfgData, &cfg); err != nil {
// 		panic(err)
// 	}

// 	appInst := app.NewApp(&cfg)
// 	log.Fatal(app.StartApplication(appInst, cfg.ListenAddr))
// }

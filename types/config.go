package types

type CrawlerConfig struct {
	// Network       string   `toml:"network"`
	// Nodes         []string `toml:"nodes"`
	// FailSleepTime string   `toml:"fail_sleep_time"`
	// PollSleepTime string   `toml:"poll_sleep_time"`
	// BlockTime     string   `toml:"block_time"`
	NumWorkers uint `toml:"num_workers"`
}

type Config struct {
	ListenAddr       string   `toml:"laddr,omitempty"`
	OnenodeAPIKey    string   `toml:"api_key"`
	OnenodeAPISecret string   `toml:"api_secret"`
	APINode          []string `toml:"api_node"`

	// Crawlers []CrawlerConfig `toml:"crawlers"`
}

package types

type Parser interface {
	Read(filename string) ([]byte, error)
	Print(v interface{})
	Parse(data []byte, v interface{}) error
	// Result()
	// GetTx(hash string) (*TxInfo, error)
	// GetTxsFromHeight(height uint64) ([]string, error)
	// GetLatestHeight() (uint64, error)
	// GetNetwork() Network
}

type Collector interface {
	GetEndpoint() string
	GetAccessKey() string
	GetParser() Parser
	GetStorage() Storage
	Get(url string, params map[string]string) ([]byte, error)
}

type Storage interface {
	CloseDB()
	Create(v interface{}) error
}

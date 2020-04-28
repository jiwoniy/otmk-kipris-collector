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

type Crawler interface {
	// GetTx(hash string) (*TxInfo, error)
	// GetTxsFromHeight(height uint64) ([]string, error)
	// GetLatestHeight() (uint64, error)
	// GetNetwork() Network
}

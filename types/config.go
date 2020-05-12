package types

type RestConfig struct {
	ListenAddr string
}

type CollectorConfig struct {
	Endpoint  string `json:"endpoint"`
	AccessKey string `json:"access_key"`
	// ListenAddr   string `json:"listen_addr"`
	DbType       string `json:"dbType"`
	DbConnString string `json:"DbConnString"`
}

type QueryConfig struct {
	DbType       string `json:"dbType"`
	DbConnString string `json:"DbConnString"`
}

type StorageConfig struct {
	// base on gorm format
	DbType       string `json:"dbType"`
	DbConnString string `json:"dbConnString"`
}

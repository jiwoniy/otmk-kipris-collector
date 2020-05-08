package types

type RestConfig struct {
	ListenAddr string
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

package types

type RestConfig struct {
	ListenAddr string
}

type ApplicationConfig struct {
	Prod CollectorConfig `json:"prod"`
	Dev  CollectorConfig `json:"dev"`
	Test CollectorConfig `json:"test"`
}

type KiprisConfig struct {
	Endpoint  string `json:"endpoint"`
	AccessKey string `json:"access_key"`
}

type DbConfig struct {
	DbType       string `json:"dbType"`
	DbConnString string `json:"dbConnString"`
}

type CollectorConfig struct {
	KiprisConfig
	DbConfig
	RestConfig
}

type QueryConfig struct {
	DbType       string `json:"dbType"`
	DbConnString string `json:"DbConnString"`
}

type StorageConfig struct {
	DbType       string `json:"dbType"`
	DbConnString string `json:"dbConnString"`
}

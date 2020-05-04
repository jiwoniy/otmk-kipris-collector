package storage

import (
	"errors"
	"fmt"
	"kipris-collector/model"
	"kipris-collector/types"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type StorageConfig struct {
	// base on gorm format
	DbType       string `json:"dbType"`
	DbConnString string `json:"dbConnString"`
}

type storage struct {
	db *gorm.DB
}

func open(dbType string, dbConnString string) (*gorm.DB, error) {
	db, err := gorm.Open(dbType, dbConnString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.TradeMarkInfo{}, &model.TrademarkDesignationGoodstInfo{})
}

func NewStorage(config StorageConfig) (types.Storage, error) {
	db, err := open(config.DbType, config.DbConnString)
	if err != nil {
		return nil, err
	}

	migrate(db)

	return &storage{
		db: db,
	}, nil
}

func (s *storage) CloseDB() {
	s.db.Close()
}

func (s *storage) Create(v types.Model) error {
	if v.Valid() == false {
		return errors.New(fmt.Sprintf("Not a valid struct %v", v))
	}
	if isCheck := s.db.NewRecord(v); isCheck == false {
		return errors.New(fmt.Sprintf("Can not create %v", v))
	}

	s.db.Create(v)

	if isFail := s.db.NewRecord(v); isFail == true {
		return errors.New(fmt.Sprintf("Fail to create data(maybe ApplicationNumber already exist) %v", v))
	}
	return nil
}

func (s *storage) GetTradeMarkInfo(v model.TradeMarkInfo, data *model.TradeMarkInfo) {
	s.db.Where(&v).First(&data)
}

func (s *storage) GetTrademarkDesignationGoodstInfo(v model.TrademarkDesignationGoodstInfo, data *model.TrademarkDesignationGoodstInfo) {
	s.db.Where(&v).First(&data)
}

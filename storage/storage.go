package storage

import (
	"errors"
	"fmt"
	"kipris-collector/model"
	"kipris-collector/types"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

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
	db.AutoMigrate(&model.TradeMarkInfo{})
}

func NewStorage(config StorageConfig) (types.Storage, error) {
	db, err := open(config.DbType, config.DbConnString)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.TradeMarkInfo{})

	return &storage{
		db: db,
	}, nil
}

func (s *storage) CloseDB() {
	s.db.Close()
}

func (s *storage) Create(v interface{}) error {
	if isCheck := s.db.NewRecord(v); isCheck == false {
		return errors.New(fmt.Sprintf("Can not create %v", v))
	}

	s.db.Create(v)

	if isFail := s.db.NewRecord(v); isFail == true {
		return errors.New(fmt.Sprintf("Fail to create data(maybe ApplicationNumber already exist) %v", v))
	}
	return nil
}

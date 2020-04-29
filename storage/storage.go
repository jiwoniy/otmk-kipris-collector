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

func migrade(db *gorm.DB) {
	db.AutoMigrate(&model.TradeMarkInfo{})
}

func NewStorage(config storageConfig) (types.Storage, error) {
	db, err := open(config.DbType, config.DbConnString)
	if err != nil {
		return nil, err
	}

	return &storage{
		db: db,
	}, nil
}

func (s *storage) CloseDB() {
	s.db.Close()
}

func (s *storage) Create(v interface{}) error {
	if isNotExist := s.db.NewRecord(v); isNotExist == false {
		return errors.New(fmt.Sprintf("This data already exist %v", v))
	}

	s.db.Create(&v)

	if isFail := s.db.NewRecord(v); isFail == true {
		return errors.New(fmt.Sprintf("create data is fail %v", v))
	}
	return nil
}

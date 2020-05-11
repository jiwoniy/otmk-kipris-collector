package storage

import (
	"errors"
	"fmt"

	"github.com/jiwoniy/otmk-kipris-collector/model"
	"github.com/jiwoniy/otmk-kipris-collector/types"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type storage struct {
	db *gorm.DB
}

func open(dbType string, dbConnString string) (*gorm.DB, error) {
	db, err := gorm.Open(dbType, dbConnString)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.TradeMarkInfo{}, &model.TrademarkDesignationGoodstInfo{}, &model.KiprisCollectorStatus{}, &model.KiprisCollectorHistory{}, &model.KiprisApplicationNumber{})
}

func NewStorage(config types.StorageConfig) (types.Storage, error) {
	db, err := open(config.DbType, config.DbConnString)
	if err != nil {
		return nil, err
	}

	migrate(db)

	return &storage{
		db: db,
	}, nil
}

func (s *storage) GetDB() *gorm.DB {
	return s.db
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

func (s *storage) GetKiprisApplicationNumber(v model.KiprisApplicationNumber, data *model.KiprisApplicationNumber) {
	s.db.Where(&v).First(&data)
}

func (s *storage) GetKiprisApplicationNumberList(v model.KiprisApplicationNumber, data *[]model.KiprisApplicationNumber, startSerialNumber int, endSerialNumber int) {
	tx := s.db.Where(&v)
	if startSerialNumber > 0 && endSerialNumber > 0 {
		tx.Find(&data, "serial_number >= ? AND serial_number <= ?", startSerialNumber, endSerialNumber)
	} else {
		tx.Find(&data)
	}
}

func (s *storage) GetKiprisCollector(v model.KiprisCollectorStatus, data *model.KiprisCollectorStatus) {
	s.db.Where(&v).Find(&data)
}

func (s *storage) GetTradeMarkInfo(v model.TradeMarkInfo, data *model.TradeMarkInfo) {
	s.db.Where(&v).First(&data)
}

func (s *storage) GetTrademarkDesignationGoodstInfo(v model.TrademarkDesignationGoodstInfo, data *[]model.TrademarkDesignationGoodstInfo) {
	s.db.Where(&v).Find(&data)
	// s.db.Where(&v).First(&data)
}

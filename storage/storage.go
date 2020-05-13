package storage

import (
	"errors"
	"fmt"

	"github.com/jiwoniy/otmk-kipris-collector/model"
	"github.com/jiwoniy/otmk-kipris-collector/pagination"
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
	// KiprisApplicationNumber ==> 이 곳에 appplication number가 등록이 되어있어야 Collector가 수행할 수 있다.
	db.AutoMigrate(&model.KiprisTask{}, &model.TradeMarkInfo{}, &model.TrademarkDesignationGoodstInfo{}, &model.KiprisCollectorStatus{}, &model.KiprisCollectorHistory{}, &model.KiprisApplicationNumber{})
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

func (s *storage) CreateTask(applicationNumbers *[]model.KiprisApplicationNumber) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		task := model.KiprisTask{}
		if err := tx.Create(&task).Error; err != nil {
			return err
		}

		tx.First(&task)

		for _, applicationNumber := range *applicationNumbers {
			kiprisApplicationNumber := model.KiprisApplicationNumber{
				ApplicationNumber: applicationNumber.ApplicationNumber,
				ProductCode:       applicationNumber.ProductCode,
				Year:              applicationNumber.Year,
				SerialNumber:      applicationNumber.SerialNumber,
				TaskId:            task.ID,
			}
			if err := tx.Create(&kiprisApplicationNumber).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *storage) GetTaskList(page int, size int) (*pagination.Paginator, error) {
	searchResult := make([]model.KiprisTask, 0)
	tx := s.db

	// TODO
	// tx = tx.Table("kipris_tasks").Where("completed IS NULL")
	// tx = tx.Table("kipris_tasks").Where("completed IS NOT NULL")
	tx = tx.Table("kipris_tasks")

	paginator := pagination.Paging(&pagination.Param{
		DB:      tx,
		Page:    page,
		Limit:   size,
		OrderBy: []string{"id asc"},
		ShowSQL: true,
	}, &searchResult)

	return paginator, nil
}

func (s *storage) GetTaskById(taskId int64) (model.KiprisTask, error) {
	currentTask := model.KiprisTask{}

	if err := s.db.Table("kipris_tasks").Where("id = ?", taskId).First(&currentTask).Error; err != nil {
		return currentTask, fmt.Errorf("[StartCrawler Task Id] %d Step 1 - Get Task", taskId)
	}
	return currentTask, nil
}

func (s *storage) GetTaskApplicationNumberList(tx *gorm.DB, taskId int64, paginationParams ...int) (*pagination.Paginator, error) {
	searchResult := make([]model.KiprisApplicationNumber, 0)
	tx.Table("kipris_application_numbers").Where("task_id = ?", taskId)

	if len(paginationParams) < 2 {
		var count int
		tx.Count(&count)
		paginatorIns := pagination.Paging(&pagination.Param{
			DB:      tx,
			Page:    1,
			Limit:   count,
			OrderBy: []string{"id asc"},
			ShowSQL: true,
		}, &searchResult)
		return paginatorIns, nil
	} else {
		paginatorIns := pagination.Paging(&pagination.Param{
			DB:      tx,
			Page:    paginationParams[0],
			Limit:   paginationParams[1],
			OrderBy: []string{"id asc"},
			ShowSQL: true,
		}, &searchResult)
		return paginatorIns, nil
	}
}

// About kipris application number
func (s *storage) GetYearLastApplicationNumber(year string) string {
	var data model.KiprisApplicationNumber
	s.db.Table("kipris_application_numbers").Where("year = ?", year).Order("application_number desc").Last(&data)
	return data.ApplicationNumber
}

// func (s *storage) GetKiprisApplicationNumber(v model.KiprisApplicationNumber, data *model.KiprisApplicationNumber) {
// 	s.db.Where(&v).First(&data)
// }

// func (s *storage) GetKiprisApplicationNumberList(v model.KiprisApplicationNumber, data *[]model.KiprisApplicationNumber, startSerialNumber int, endSerialNumber int, page int, size int) (*pagination.Paginator, error) {
// 	tx := s.db

// 	if startSerialNumber > 0 && endSerialNumber > 0 {
// 		tx = tx.Where(&v).Where("serial_number >= ? AND serial_number <= ?", startSerialNumber, endSerialNumber)
// 	} else {
// 		tx = tx.Where(&v)
// 	}

// 	paginator := pagination.Paging(&pagination.Param{
// 		DB:      tx,
// 		Page:    page,
// 		Limit:   size,
// 		OrderBy: []string{"application_number asc"},
// 		ShowSQL: true,
// 	}, data)

// 	return paginator, nil
// }

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

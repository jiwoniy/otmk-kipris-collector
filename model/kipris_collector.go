package model

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type KiprisTask struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Started   time.Time
	Completed time.Time
}

// kipirs aplication number list
type KiprisApplicationNumber struct {
	gorm.Model
	TaskId            int64  `gorm:"not null;index;" validate:"required"`
	ApplicationNumber string `gorm:"not null;index;" validate:"required"`
	ProductCode       string `gorm:"not null;index" validate:"required"`
	Year              string `gorm:"not null;index" validate:"required"`
	SerialNumber      int    `gorm:"not null;index" validate:"required"`
	isExist           bool
}

type KiprisCollectorStatus struct {
	gorm.Model
	ApplicationNumber                  string               `gorm:"not null;" validate:"required"`
	TradeMarkInfoStatus                KiprisResponseStatus `validate:"required"`
	TradeMarkDesignationGoodInfoStatus KiprisResponseStatus `validate:"required"`
}

type KiprisCollectorHistory struct {
	gorm.Model
	ApplicationNumber string `gorm:"not null;" validate:"required"`
	IsSuccess         bool
	Error             string
}

func (data *KiprisApplicationNumber) Valid() bool {
	return true
}

func (data *KiprisCollectorStatus) Valid() bool {
	validate = validator.New()
	err := validate.Struct(data)

	if err != nil {
		// from here you can create your own error messages in whatever language you wish
		return false
	}
	return true
}

func (data *KiprisCollectorHistory) Valid() bool {
	validate = validator.New()
	err := validate.Struct(data)

	if err != nil {
		// from here you can create your own error messages in whatever language you wish
		return false
	}
	return true
}

package model

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type KiprisTask struct {
	ID          int64      `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt   time.Time  `json:"createdAt,omitempty"`
	UpdatedAt   time.Time  `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time `sql:"index" json:"-"`
	StartedAt   time.Time  `json:"startedAt,omitempty"`
	CompletedAt time.Time  `json:"completedAt,omitempty"`
}

// kipirs aplication number list
type KiprisApplicationNumber struct {
	gorm.Model
	TaskId            int64  `gorm:"not null;" validate:"required"  json:"taskId,omitempty"`
	ApplicationNumber string `gorm:"not null;" validate:"required" json:"application_number,omitempty"`
	ProductCode       string `gorm:"not null;" validate:"required" json:"product_code,omitempty"`
	Year              string `gorm:"not null;" validate:"required" json:"year,omitempty"`
	SerialNumber      int    `gorm:"not null;" validate:"required" json:"serial_number,omitempty"`
	isExist           bool   `json:"_"`
}

type KiprisCollectorStatus struct {
	gorm.Model
	TaskId                             int64                `gorm:"not null;index;" validate:"required"`
	ApplicationNumber                  string               `gorm:"not null;" validate:"required"`
	TradeMarkInfoStatus                KiprisResponseStatus `validate:"required"`
	TradeMarkDesignationGoodInfoStatus KiprisResponseStatus `validate:"required"`
}

type KiprisCollectorHistory struct {
	gorm.Model
	ApplicationNumber string `gorm:"not null;" validate:"required"`
	TaskId            int64  `gorm:"not null;index;" validate:"required"`
	IsSuccess         bool
	Error             string `gorm:"type:text"`
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

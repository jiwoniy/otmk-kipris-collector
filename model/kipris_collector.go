package model

import (
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

// kipirs aplication number list
type KiprisApplicationNumber struct {
	gorm.Model
	ApplicationNumber string `gorm:"not null;unique;" validate:"required"`
	ProductCode       string `gorm:"not null;" validate:"required"`
	Year              string `gorm:"not null;" validate:"required"`
	SerialNumber      int    `gorm:"not null;" validate:"required"`
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

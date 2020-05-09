package model

import (
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type KiprisApplicationNumber struct {
	gorm.Model
	ApplicationNumber string `gorm:"not null;" validate:"required"`
}

type KiprisCollector struct {
	gorm.Model
	ApplicationNumber                  string               `gorm:"not null;" validate:"required"`
	TradeMarkInfoStatus                KiprisResponseStatus `validate:"required"`
	TradeMarkDesignationGoodInfoStatus KiprisResponseStatus `validate:"required"`
	// Error             string
}

type KiprisCollectorHistory struct {
	gorm.Model
	ApplicationNumber string `gorm:"not null;" validate:"required"`
	IsSuccess         bool   `validate:"required"`
	Error             string `validate:"required"`
	// Error             string
}

func (data *KiprisApplicationNumber) Valid() bool {
	return true
}

func (data *KiprisCollector) Valid() bool {
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

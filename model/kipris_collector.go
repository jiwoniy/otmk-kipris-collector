package model

import (
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type KiprisCollector struct {
	gorm.Model
	ApplicationNumber                  string               `gorm:"unique;not null;" validate:"required"`
	TradeMarkInfoStatus                KiprisResponseStatus `validate:"required"`
	TradeMarkDesignationGoodInfoStatus KiprisResponseStatus `validate:"required"`
	// Error             string
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

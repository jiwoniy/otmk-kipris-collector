package model

import (
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

// `TrademarkDesignationGoodstInfo` belongs to `TradmarkInfo`is the foreign key
type TrademarkDesignationGoodstInfo struct {
	gorm.Model
	// XMLName                                       xml.Name   `xml:"trademarkDesignationGoodstInfo"`
	ApplicationNumber                             string     `gorm:"association_foreignkey:ApplicationNumber;not null;" validate:"required"` // use Refer as association foreign ke
	DesignationGoodsSerialNumber                  TrimString `xml:"DesignationGoodsSerialNumber,omitempty" gorm:"not null;index;" validate:"required"`
	DesignationGoodsClassificationInformationCode TrimString `xml:"DesignationGoodsClassificationInformationCode,omitempty" gorm:"not null;index;" validate:"required"`
	SimilargroupCode                              TrimString `xml:"SimilargroupCode,omitempty" gorm:"not null;index;" validate:"required"`
	DesignationGoodsHangeulName                   TrimString `xml:"DesignationGoodsHangeulName,omitempty" gorm:"not null;index;" validate:"required"`
	DesignationGoodsEnglishsentenceName           TrimString `xml:"DesignationGoodsEnglishsentenceName,omitempty"`
}

func (data *TrademarkDesignationGoodstInfo) Valid() bool {
	validate = validator.New()
	err := validate.Struct(data)

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		// if _, ok := err.(*validator.InvalidValidationError); ok {
		// 	// fmt.Println(err)
		// 	return false
		// }

		// for _, err := range err.(validator.ValidationErrors) {
		// 	fmt.Println(err.Namespace())
		// 	fmt.Println(err.Field())
		// 	fmt.Println(err.StructNamespace())
		// 	fmt.Println(err.StructField())
		// 	fmt.Println(err.Tag())
		// 	fmt.Println(err.ActualTag())
		// 	fmt.Println(err.Kind())
		// 	fmt.Println(err.Type())
		// 	fmt.Println(err.Value())
		// 	fmt.Println(err.Param())
		// }

		// from here you can create your own error messages in whatever language you wish
		return false
	}
	return true
}

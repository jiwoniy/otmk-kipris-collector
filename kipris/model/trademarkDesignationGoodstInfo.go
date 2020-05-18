package model

import (
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

// `TrademarkDesignationGoodstInfo` belongs to `TradmarkInfo`is the foreign key
type TrademarkDesignationGoodstInfo struct {
	gorm.Model
	// XMLName                                       xml.Name   `xml:"trademarkDesignationGoodstInfo"`
	ApplicationNumber                             string     `gorm:"association_foreignkey:ApplicationNumber;not null;" validate:"required" json:"-"` // use Refer as association foreign ke
	DesignationGoodsSerialNumber                  TrimString `xml:"DesignationGoodsSerialNumber,omitempty" gorm:"not null;index;" validate:"required" json:"designationGoodsSerialNumber"`
	DesignationGoodsClassificationInformationCode TrimString `xml:"DesignationGoodsClassificationInformationCode,omitempty" gorm:"not null;index;" validate:"required" json:"designationGoodsClassificationInformationCode"`
	SimilargroupCode                              TrimString `xml:"SimilargroupCode,omitempty" gorm:"not null;index;" validate:"required" json:"similargroupCode"`
	DesignationGoodsHangeulName                   TrimString `xml:"DesignationGoodsHangeulName,omitempty" gorm:"type:text;" validate:"required" json:"designationGoodsHangeulName"`
	DesignationGoodsEnglishsentenceName           TrimString `xml:"DesignationGoodsEnglishsentenceName,omitempty" json:"designationGoodsEnglishsentenceName"`
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

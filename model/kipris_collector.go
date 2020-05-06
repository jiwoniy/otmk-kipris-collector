package model

import (
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type KiprisCollector struct {
	gorm.Model
	ApplicationNumber string               `gorm:"unique;not null;" validate:"required"`
	Status            KiprisResponseStatus `validate:"required"`
	// Error             string
}

func (data *KiprisCollector) Valid() bool {
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

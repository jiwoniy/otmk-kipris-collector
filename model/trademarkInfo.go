package model

import (
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type TradeMarkInfo struct {
	// ID        int64 `gorm:"primary_key,AUTO_INCREMENT"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt *time.Time
	gorm.Model

	SerialNumber       TrimString `xml:"SerialNumber"`
	ApplicationNumber  string     `xml:"ApplicationNumber" gorm:"unique;not null" validate:"required"`
	AppReferenceNumber TrimString `xml:"AppReferenceNumber"`
	ApplicationDate    TrimString `xml:"ApplicationDate"`
	PublicNumber       TrimString `xml:"PublicNumber"`

	PublicDate               TrimString `xml:"PublicDate"`
	RegistrationPublicNumber TrimString `xml:"RegistrationPublicNumber"`
	RegistrationPublicDate   TrimString `xml:"RegistrationPublicDate"`
	RegistrationNumber       TrimString `xml:"RegistrationNumber"`
	RegReferenceNumber       TrimString `xml:"RegReferenceNumber"`

	RegistrationDate       TrimString `xml:"RegistrationDate"`
	PriorityClaimNumber    TrimString `xml:"PriorityClaimNumber"`
	PriorityClaimDate      TrimString `xml:"PriorityClaimDate"`
	ApplicationStatus      TrimString `xml:"ApplicationStatus"`
	GoodClassificationCode TrimString `xml:"GoodClassificationCode"`

	ViennaCode                  TrimString `xml:"ViennaCode"`
	ApplicantName               TrimString `xml:"ApplicantName"`
	AgentName                   TrimString `xml:"AgentName"`
	RegistrationRightholderName TrimString `xml:"RegistrationRightholderName"`
	Title                       TrimString `xml:"Title"`

	FulltextExistFlag TrimString `xml:"FulltextExistFlag"`
	ImagePath         TrimString `xml:"ImagePath"`
	ThumbnailPath     TrimString `xml:"ThumbnailPath"`
}

func (data *TradeMarkInfo) Valid() bool {
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

type Items struct {
	// XMLName                        xml.Name                         `xml:"items"`
	TrademarkDesignationGoodstInfo []TrademarkDesignationGoodstInfo `xml:"trademarkDesignationGoodstInfo"`
	TradeMarkInfo                  TradeMarkInfo                    `xml:"TradeMarkInfo"`
	TotalSearchCount               TotalSearchCount                 `xml:"TotalSearchCount"`
}

type TotalSearchCount TrimString

type Body struct {
	// XMLName xml.Name `xml:"body"`
	Items Items `xml:"items"`
}

type Header struct {
	// XMLName    xml.Name `xml:"header"`
	ResultCode TrimString `xml:"resultCode"`
	ResultMsg  TrimString `xml:"resultMsg"`
}

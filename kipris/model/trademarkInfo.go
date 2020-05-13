package model

// https://m.blog.naver.com/PostView.nhn?blogId=s0215hc&logNo=221475878207&proxyReferer=https:%2F%2Fwww.google.com%2F
// https://blog.lael.be/post/917
import (
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type TradeMarkInfo struct {
	gorm.Model

	SerialNumber TrimString `xml:"SerialNumber" json:"serialNumber"`

	ApplicationNumber  string     `xml:"ApplicationNumber" json:"applicationNumber" gorm:"application_number;unique;not null" validate:"required"`
	AppReferenceNumber TrimString `xml:"AppReferenceNumber" json:"appReferenceNumber" gorm:"app_reference_number"`
	ApplicationDate    TrimString `xml:"ApplicationDate" json:"applicationDate"`
	PublicNumber       TrimString `xml:"PublicNumber" json:"publicNumber"`

	PublicDate               TrimString `xml:"PublicDate" json:"publicDate"`
	RegistrationPublicNumber TrimString `xml:"RegistrationPublicNumber" json:"registrationPublicNumber"`
	RegistrationPublicDate   TrimString `xml:"RegistrationPublicDate" json:"registrationPublicDate"`
	RegistrationNumber       TrimString `xml:"RegistrationNumber" json:"registrationNumber"`
	RegReferenceNumber       TrimString `xml:"RegReferenceNumber" json:"regReferenceNumber"`

	RegistrationDate       TrimString `xml:"RegistrationDate" json:"registrationDate"`
	PriorityClaimNumber    TrimString `xml:"PriorityClaimNumber" json:"priorityClaimNumber"`
	PriorityClaimDate      TrimString `xml:"PriorityClaimDate" json:"priorityClaimDate"`
	ApplicationStatus      TrimString `xml:"ApplicationStatus" json:"applicationStatus"`
	GoodClassificationCode TrimString `xml:"GoodClassificationCode" json:"goodClassificationCode"`

	ViennaCode                  TrimString `xml:"ViennaCode" json:"viennaCode"`
	ApplicantName               TrimString `xml:"ApplicantName" json:"applicantName"`
	AgentName                   TrimString `xml:"AgentName" json:"agentName"`
	RegistrationRightholderName TrimString `xml:"RegistrationRightholderName" json:"registrationRightholderName"`
	Title                       TrimString `xml:"Title" json:"title"`

	FulltextExistFlag TrimString `xml:"FulltextExistFlag" json:"fulltextExistFlag"`
	ImagePath         TrimString `xml:"ImagePath" json:"imagePath"`
	ThumbnailPath     TrimString `xml:"ThumbnailPath" json:"thumbnailPath"`

	TrademarkDesignationGoodstInfos []TrademarkDesignationGoodstInfo `json:"trademarkDesignationGoodstInfos"`
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

type TotalSearchCount TrimString

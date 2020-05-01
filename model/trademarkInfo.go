package model

import "time"

type TradeMarkInfo struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	// XMLName            xml.Name   `xml:"TradeMarkInfo"`
	SerialNumber       TrimString `xml:"SerialNumber"`
	ApplicationNumber  TrimString `xml:"ApplicationNumber" gorm:"unique;not null"`
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

type TrademarkDesignationGoodstInfo struct {
	// XMLName                                       xml.Name   `xml:"trademarkDesignationGoodstInfo"`
	DesignationGoodsSerialNumber                  TrimString `xml:"DesignationGoodsSerialNumber,omitempty"`
	DesignationGoodsClassificationInformationCode TrimString `xml:"DesignationGoodsClassificationInformationCode,omitempty"`
	SimilargroupCode                              TrimString `xml:"SimilargroupCode,omitempty"`
	DesignationGoodsHangeulName                   TrimString `xml:"DesignationGoodsHangeulName,omitempty"`
	DesignationGoodsEnglishsentenceName           TrimString `xml:"DesignationGoodsEnglishsentenceName,omitempty"`
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

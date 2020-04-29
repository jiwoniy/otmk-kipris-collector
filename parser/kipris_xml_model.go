package parser

type TradeMarkInfo struct {
	// XMLName            xml.Name   `xml:"TradeMarkInfo"`
	SerialNumber       TrimString `xml:"SerialNumber"`
	ApplicationNumber  TrimString `xml:"ApplicationNumber"`
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
	TrademarkDesignationGoodstInfo []TrademarkDesignationGoodstInfo `xml:"trademarkDesignationGoodstInfo,omitempty"`
	TradeMarkInfo                  TradeMarkInfo                    `xml:"TradeMarkInfo,omitempty"`
	TotalSearchCount               TrimString                       `xml:"TotalSearchCount,omitempty"`
}

type Body struct {
	// XMLName xml.Name `xml:"body"`
	Items Items `xml:"items"`
}

type Header struct {
	// XMLName    xml.Name `xml:"header"`
	ResultCode TrimString `xml:"resultCode"`
	ResultMsg  TrimString `xml:"resultMsg"`
}

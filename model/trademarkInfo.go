package model

import "time"

// sync with parser/kipris_xml_model
type TradeMarkInfo struct {
	ApplicationNumber string `gorm:"primary_key"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time

	SerialNumber       string
	AppReferenceNumber string
	ApplicationDate    string
	PublicNumber       string

	PublicDate               string
	RegistrationPublicNumber string
	RegistrationPublicDate   string
	RegistrationNumber       string
	RegReferenceNumber       string

	RegistrationDate       string
	PriorityClaimNumber    string
	PriorityClaimDate      string
	ApplicationStatus      string
	GoodClassificationCode string

	ViennaCode                  string
	ApplicantName               string
	AgentName                   string
	RegistrationRightholderName string
	Title                       string

	FulltextExistFlag string
	ImagePath         string
	ThumbnailPath     string
}

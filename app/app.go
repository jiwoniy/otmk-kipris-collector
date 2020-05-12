package app

import (
	"github.com/jiwoniy/otmk-kipris-collector/types"
)

type Application struct {
	collector types.RestClient
}

var app Application

func NewApplication(collector types.RestClient) *Application {
	app.collector = collector
	return &app
}

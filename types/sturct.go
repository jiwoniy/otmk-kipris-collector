package types

import (
	"github.com/gin-gonic/gin"
)

type RestHandler func(ctx *gin.Context)

type RestMethod struct {
	Path    string
	Handler RestHandler
}

type RestFailResponse struct {
	Error string
}

type TaskParameters struct {
	ProductCode       string
	Year              string
	SerialNumberRange string
	Page              int
	Size              int
}

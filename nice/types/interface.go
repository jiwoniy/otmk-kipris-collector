package types

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jiwoniy/otmk-kipris-collector/nice/schema"
	"github.com/jiwoniy/otmk-kipris-collector/utils"
)

type RestHandler func(ctx *gin.Context)

type RestMethod struct {
	Path    string
	Handler RestHandler
}

type Result struct {
	Data interface{} `json:"data,omitempty"`
}

type Storage struct {
	DB *gorm.DB
}

type ImportClient interface {
	GetStorage() *Storage
	ImportNiceCsv(folderPath string, db *gorm.DB)
	GetNiceList(result *[]schema.NiceClassification)
}

type Client interface {
	GetStorage() *Storage
	GetMethods() ([]RestMethod, error)
}

type QueryClient interface {
	Client
	SearchSimilarGroups(text string, classificationCode string, size int, page int) (*utils.Paginator, error)
	GetSimilarCodeText(id string) (string, error)
}

package query

import (
	"github.com/jinzhu/gorm"
	"github.com/jiwoniy/otmk-kipris-collector/nice/types"
)

// type storage struct {
// 	db *gorm.DB
// }

func NewStorage(db *gorm.DB) *types.Storage {
	// db.AutoMigrate(&sdk.AccountTag{}, &sdk.AccountModel{}, &sdk.KeyLocation{})
	return &types.Storage{
		Db: db,
	}
}

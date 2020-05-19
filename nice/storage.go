package nice

import (
	"github.com/jinzhu/gorm"
	"github.com/jiwoniy/otmk-kipris-collector/nice/types"
)

func NewStorage(db *gorm.DB) *types.Storage {
	return &types.Storage{
		DB: db,
	}
}

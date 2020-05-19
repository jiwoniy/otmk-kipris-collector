package schema

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 상품/서비스업 항목 류 추천
type RecommendClassfication struct {
	Id        uint64   `json:"id"`
	NameKo    string   `json:"name_ko"`
	NameEn    string   `json:"name_en"`
	Recommend []string `json:"recommend"`
	Related   []string `json:"related"`
}

type commonModelFields struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type NiceClassification struct {
	ID                 uint `json:"id"`
	ClassificationCode string
	Code               string
	Name_EN            string
	Name_KO            string
	Source             string
}

type SimilarCodeGroupReciver struct {
	ClassificationCode string `json:"classificationCode"`
	Code               string `json:"code"`
	Count              int    `json:"count"`
}

type SimilarCodeGroup struct {
	ID uint `json:"id"`
	SimilarCodeGroupReciver
	Summary_EN string `json:"summaryEn"`
	Summary_KO string `json:"summaryKo"`
}

type ProductDomain struct {
	gorm.Model
	Name_EN   string
	Name_KO   string
	Recommend []string
	Related   []string
}

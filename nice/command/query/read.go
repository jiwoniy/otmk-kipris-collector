package query

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jiwoniy/otmk-kipris-collector/nice/schema"
	"github.com/jiwoniy/otmk-kipris-collector/nice/types"
	"github.com/jiwoniy/otmk-kipris-collector/utils"
)

func init() {}

type keeper struct {
	storage *types.Storage
}

type response struct {
	Data interface{} `json:"data"`
}

type similarityGroupsPaginatorResponse struct {
	SimilarCodeGroup   interface{} `json:"similarCodeGroup"`
	NiceClassification interface{} `json:"niceClassfication"`
	TotalRecord        int         `json:"total_record"`
	TotalPage          int         `json:"total_page"`
	Page               int         `json:"page"`
}

func NewKeeper(db *gorm.DB) types.QueryClient {
	return &keeper{
		storage: NewStorage(db),
	}
}

func (query *keeper) GetStorage() *types.Storage {
	return query.storage
}

func (query *keeper) GetMethods() ([]types.RestMethod, error) {
	restMethods := make([]types.RestMethod, 1)
	restMethods = append(restMethods, types.RestMethod{
		Path:    "/search",
		Handler: query.getSearch,
	},
		types.RestMethod{
			Path:    "/searchText",
			Handler: query.getSearchText,
		},
	)
	return restMethods, nil
}

// func writeResponse(ctx *gin.Context, data interface{}) {
// 	var response Result = Result{
// 		Data: data,
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }

func (query *keeper) getSearch(ctx *gin.Context) {
	text := ctx.Query("text")
	classificationCode := ctx.Query("classificationCode")
	page := ctx.DefaultQuery("page", "1")
	size := ctx.DefaultQuery("size", "20")

	sizeInt, _ := strconv.Atoi(size)
	pageInt, _ := strconv.Atoi(page)

	paginatorSimilarGroups, err := query.SearchSimilarGroups(text, classificationCode, sizeInt, pageInt)
	if err != nil {
		panic(err)
	}

	niceGroups, err := query.SearchNiceClassfications(text, classificationCode)
	if err != nil {
		panic(err)
	}

	var response similarityGroupsPaginatorResponse = similarityGroupsPaginatorResponse{
		SimilarCodeGroup:   paginatorSimilarGroups.Data,
		TotalRecord:        paginatorSimilarGroups.TotalRecord,
		TotalPage:          paginatorSimilarGroups.TotalPage,
		Page:               paginatorSimilarGroups.Page,
		NiceClassification: niceGroups,
	}
	ctx.JSON(http.StatusOK, response)
}

func (query *keeper) getSearchText(ctx *gin.Context) {
	id := ctx.Query("id")
	// ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
	// ctx.Header("Access-Control-Allow-Credentials", "true")
	// ctx.Header("Access-Control-Allow-Origin", "*") // 에러 - Credentials이 true일 경우 특정 URL 1개만 허용
	// ctx.Header("Access-Control-Allow-Methods", "GET")
	// ctx.Next()

	text, err := query.GetSimilarCodeText(id)
	if err != nil {
		panic(err)
	}

	var response response = response{
		Data: text,
	}

	ctx.JSON(http.StatusOK, response)
}

func (query *keeper) SearchNiceClassfications(text string, classificationCode string) (*[]schema.SimilarCodeGroup, error) {
	var codesGroups []schema.SimilarCodeGroup
	tx := query.storage.Db.Table("similar_code_groups").Select("id, classification_code").Group("classification_code")

	tx = tx.Where("summary_ko LIKE ?", fmt.Sprint("%", text, "%"))
	if len(classificationCode) > 0 {
		tx = tx.Where("classification_code = ?", classificationCode)
	}

	tx.Find(&codesGroups)

	return &codesGroups, nil
}

func (query *keeper) SearchSimilarGroups(text string, classificationCode string, size int, page int) (*utils.Paginator, error) {
	var codesGroups []schema.SimilarCodeGroup
	tx := query.storage.Db.Table("similar_code_groups")
	// .Select(`id, classification_code, code, TRIM(substr(summary_ko, INSTR(summary_ko, ${휴대폰'), LENGTH(summary_ko) - INSTR(summary_ko, '휴대폰'))) as summary_ko)`)

	if len(text) > 0 {
		summaryQuery := fmt.Sprint("INSTR(summary_ko, '", text, "')")
		selectQuery := fmt.Sprint("id, classification_code, code, TRIM(substr(summary_ko, ", summaryQuery, ", LENGTH(summary_ko) - 100)) as summary_ko")
		tx = tx.Select(selectQuery)
	} else {
		tx = tx.Select("id, classification_code, code, SUBSTR(summary_ko, 0, 100) as summary_ko")
	}

	// 	select
	// summary_ko,
	// --CHARINDEX('t', 'Customer') AS MatchPosition
	// INSTR(summary_ko, '휴대폰'),
	// LENGTH(summary_ko),
	// TRIM(substr(summary_ko, INSTR(summary_ko, '휴대폰'), LENGTH(summary_ko) - INSTR(summary_ko, '휴대폰')))
	// from similar_code_groups where summary_ko like '%휴대폰%'

	tx = tx.Where("summary_ko LIKE ?", fmt.Sprint("%", text, "%"))
	if len(classificationCode) > 0 {
		tx = tx.Where("classification_code = ?", classificationCode)
	}

	paginator := utils.Paging(&utils.PaginatorParam{
		DB:      tx,
		Page:    page,
		Limit:   size,
		OrderBy: []string{"classification_code asc", "code asc"},
		ShowSQL: true,
	}, &codesGroups)

	return paginator, nil
}

// https://www.sqlitetutorial.net/sqlite-string-functions/
// select
// summary_ko,
// --CHARINDEX('t', 'Customer') AS MatchPosition
// INSTR(summary_ko, '휴대폰'),
// substr(summary_ko, INSTR(summary_ko, '휴대폰'), 30)
// from similar_code_groups where summary_ko like '%휴대폰%'

func (query *keeper) GetSimilarCodeText(id string) (string, error) {
	var codesGroup schema.SimilarCodeGroup
	tx := query.storage.Db.Table("similar_code_groups").Select("id, summary_ko")

	tx = tx.Where("id = ?", id).First(&codesGroup)

	result := codesGroup
	return result.Summary_KO, nil
}

// type Read struct{}

// func readJson() []schema.RecommendClassfication {
// 	dataPath := "./data"
// 	jsonFile, _ := os.Open(dataPath + "/recommend_classfiication.json")

// 	var datas []schema.RecommendClassfication
// 	byteValue, _ := ioutil.ReadAll(jsonFile)
// 	json.Unmarshal(byteValue, &datas)

// 	return datas
// }

// func (e *Read) GetRecommendList() map[uint64]schema.RecommendClassfication {
// 	list := readJson()

// 	m := make(map[uint64]schema.RecommendClassfication)

// 	for i := 0; i < len(list); i++ {
// 		m[list[i].Id] = list[i]
// 	}

// 	return m
// }

// func (e *Read) Search(db *gorm.DB, search map[string]string) *pagination.Paginator {
// 	var codes []schema.NiceClassification

// 	tx := db

// 	if len(search["name"]) > 0 {
// 		tx = tx.Where("name_ko LIKE ?", fmt.Sprint("%", search["name"], "%"))
// 	}

// 	if len(search["classification_code"]) > 0 {
// 		tx = tx.Where("classification_code = ?", search["classification_code"])
// 	}

// 	if len(search["code"]) > 0 {
// 		tx = tx.Where("code = ?", search["code"])
// 	}

// 	result := pagination.Paging(&pagination.Param{
// 		DB:      tx,
// 		Page:    1,
// 		Limit:   20,
// 		OrderBy: []string{"id desc"},
// 		ShowSQL: true,
// 	}, &codes)

// 	return result
// }

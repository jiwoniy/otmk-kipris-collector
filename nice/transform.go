package command

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/jiwoniy/otmk-kipris-collector/nice/schema"
)

// func createSimilarCodeGroup(db *gorm.DB, result *schema.SimilarCodeGroup) {
// 	fmt.Println(result)
// 	db.Create(&result)
// }

var (
	logBuf bytes.Buffer
	logger = log.New(&logBuf, "transform logger: ", log.Lshortfile)
)

// stream 으로 만들면 좋을텐데..
func selectSimilarCodeGroup(db *gorm.DB) {
	db.DropTable(&schema.SimilarCodeGroup{})
	db.AutoMigrate(&schema.SimilarCodeGroup{})

	var codes []schema.SimilarCodeGroupReciver
	db.Table("nice_classifications").Select("classification_code, code, COUNT(code) as count").Group("code").Find(&codes)

	var summary []string
	var summaryKo []string
	for _, code := range codes {
		// db.Table("nice_classifications").Where("code = ? AND classification_code >= ?", code.Code, code.ClassificationCode).Limit(20).Pluck("name_ko", &summary)
		db.Table("nice_classifications").Where("code = ? AND classification_code >= ?", code.Code, code.ClassificationCode).Pluck("name_ko", &summary)

		// logger.Print(fmt.Sprintf("classfication: %s -- similrar code: %s -- name: %s", code.ClassificationCode, code.Code, strings.Join(summary[:], ",")))
		summaryKo = append(summaryKo, strings.Join(summary[:], ","))
	}

	for index, code := range codes {
		data := schema.SimilarCodeGroup{
			SimilarCodeGroupReciver: code,
			Summary_KO:              summaryKo[index],
		}
		db.Create(&data)
	}
}

func Main(db *gorm.DB) {
	logFile, err := os.Create("./log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()
	logger.SetOutput(logFile)

	// log.Println("Logggggg") log.Fatalln("Fatal Error")

	selectSimilarCodeGroup(db)
}

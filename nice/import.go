package nice

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/jiwoniy/otmk-kipris-collector/nice/schema"
	"github.com/jiwoniy/otmk-kipris-collector/nice/types"
)

func init() {}

type keeper struct {
	storage *types.Storage
}

func NewKeeper(db *gorm.DB) types.ImportClient {
	return &keeper{
		storage: NewStorage(db),
	}
}

func (command *keeper) GetStorage() *types.Storage {
	return command.storage
}

func (command *keeper) ImportNiceCsv(folderPath string, db *gorm.DB) {
	migrate(db)

	files, err := getFiles(folderPath)
	if err != nil {
		panic(err)
	} else {
		for _, file := range files {
			filePath := folderPath + "/" + file.Name()
			if err := importCsvFile(filePath, db); err != nil {
				panic(err)
			}
		}
	}
}

func (command *keeper) GetNiceList(result *[]schema.NiceClassification) {
	// var count int
	storage := command.GetStorage()
	// storage.DB.Find(&data).Count(&count)
	storage.DB.Find(result)
}

// func (command *keeper) GetMethods() ([]types.RestMethod, error) {
// 	restMethods := []types.RestMethod{
// 		types.RestMethod{},
// 	}
// 	return restMethods, nil
// }

// func (command *keeper) Create() error {
// 	return nil
// }

// Ref
// https://medium.com/@barunthapa/working-with-csv-in-go-50a4f540e623
func getKey(index int) (string, error) {
	// 지정상품(국문),NICE분류,유사군코드,지정상품(영문),출처
	// j = 0 지정상품(국문)
	// j = 1 NICE분류
	// j = 2 유사군코드
	// j = 3 지정상품(영문)
	// j = 4 출처
	switch index {
	case 0:
		return "Name_KO", nil
	case 1:
		return "ClassificationCode", nil
	case 2:
		return "Code", nil
	case 3:
		return "Name_EN", nil
	case 4:
		return "Source", nil
	default:
		return "", fmt.Errorf("Not defined Key")
	}
}

func getFiles(path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return files, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&schema.NiceClassification{})
}

// func importCsvFile(db *gorm.DB, path string) error {
func importCsvFile(path string, db *gorm.DB) error {
	file, err := os.Open(path)
	if err == nil {
		reader := csv.NewReader(bufio.NewReader(file))
		rows, _ := reader.ReadAll()

		for rowIndex, row := range rows {
			if rowIndex > 0 {
				rowData := make(map[string]string)
				for columnIndex, value := range row {
					key, err := getKey(columnIndex)
					if err != nil {
						return err
					}
					rowData[key] = value
				}
				dataRow := schema.NiceClassification{
					ClassificationCode: rowData["ClassificationCode"],
					Code:               rowData["Code"],
					Name_EN:            rowData["Name_EN"],
					Name_KO:            rowData["Name_KO"],
					Source:             rowData["Source"],
				}
				db.Create(&dataRow)
			}
		}
		return nil
	}

	return err
}

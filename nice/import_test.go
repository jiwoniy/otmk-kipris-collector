package nice

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jiwoniy/otmk-kipris-collector/nice/schema"
)

func TestNiceImport(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	// db, err := gorm.Open("mysql", "nice_admin:Nice0518!@@(61.97.187.142:3306)/nice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	client := NewKeeper(db)
	// fmt.Println(client)
	// fmt.Println(client.GetStorage())
	folderPath := "./csv_test"
	client.ImportNiceCsv(folderPath, db)

	var result []schema.NiceClassification
	client.GetNiceList(&result)

	if len(result) != 17 {
		t.Error("Check data length")
	}
}

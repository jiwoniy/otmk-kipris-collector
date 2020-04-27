package saver

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(db)
	fmt.Println(err)
	defer db.Close()
}

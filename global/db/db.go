package db

import (
	global "blog/global/setting"
	"blog/internal/model/database"
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	DBEngine *gorm.DB
)

func init() {
	var err error
	DBEngine, err = database.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("db connected")
	}
}

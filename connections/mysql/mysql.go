package mysql

import (
	"fmt"
	"food_mlem/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	mysqlDb0 := conf.MYSQL_CONFIG["ENDPOINT"]
	db, err := gorm.Open(mysql.Open(mysqlDb0), &gorm.Config{})
	if err != nil {
		panic(err)
		//logs.Error(err)
		fmt.Println("CONNECT MYSQL FAILED!")
		fmt.Println(err.Error())
	}
	DB = db
	fmt.Println("CONNECT MYSQL SUCCESS!")
}

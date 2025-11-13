package Task3

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Entry() {
	var db *gorm.DB = InitDB()
	StudentRun(db)

}

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	return db
}

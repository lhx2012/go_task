package Task3

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Entry() {
	//var db *gorm.DB = InitDB()
	////StudentRun(db)
	//BankRun(db)
	//Sqlx1Run()
	//sqlxRun2()
	blogRun()
}

func InitDB() *gorm.DB {
	//db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	//if err != nil {
	//	panic("连接数据库失败：" + err.Error())
	//}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}

	return db
}

func closeDB(db *gorm.DB) {
	// 记得在合适的时候关闭数据库连接
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}

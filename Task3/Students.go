package Task3

import (
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	ID    int32
	Name  string
	Age   uint8
	Grade string
}

func StudentRun(db *gorm.DB) {
	//新建students表
	db.AutoMigrate(&Student{})

	//queryName := Student{}
	//db.Table("students").Where("name ==", "张三")(&queryName)
	//if len() == 0 {
	//	//插入一条数据
	//	student := Student{Name: "张三", Age: 18, Grade: "三年级"}
	//	db.Create(&student)
	//}

	//查询年龄大于18的数据
	ageCondition := Student{}

	//db.Table("students").Where("age > ?", student.Age).Scan(&ageCondition)

	db.Table("students").Scopes(AgeFilter).Scan(&ageCondition)
	fmt.Println(ageCondition)
}

func AgeFilter(db *gorm.DB) *gorm.DB {
	return db.Where("age >= ?", "18")
}

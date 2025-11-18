package Task3

import (
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	ID    int32 `gorm:"primaryKey"`
	Name  string
	Age   uint8
	Grade string
}

func StudentRun(db *gorm.DB) {
	//新建students表
	err := db.AutoMigrate(&Student{})
	if err != nil {
		return
	}

	//插入一条数据
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	db.Create(&student)

	//查询年龄大于18的数据
	var ageCondition []Student
	//db.Table("students").Scopes(AgeFilter).Scan(&ageCondition)

	db.Model(&ageCondition).Where("age >= ?", "18").Find(&ageCondition)
	for _, value := range ageCondition {
		fmt.Printf("查找到数据:%v\n", value)
	}

	//更新张三的年级为四年级
	var std Student
	db.Debug().Model(&std).Where("Name==?", "张三").Update("Grade", "四年级")
	//db.Model(&std).Updates(Student{Grade: "四年级", Name: "张三"})
	//db.Model(&std).Updates(map[string]interface{}{"Grade": "四年级", "Name": "张三"})

	//删除年龄小于15的学生
	db.Debug().Delete(&std, "age < ?", "25")
}

func AgeFilter(db *gorm.DB) *gorm.DB {
	return db.Where("age >= ?", "18")
}

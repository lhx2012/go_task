package Task3

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//题目1：使用SQL扩展库进行查询
//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
//要求 ：
//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

func insertData(db *sqlx.DB) error {
	createSql := `
	CREATE TABLE IF NOT EXISTS employees(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    department TEXT NOT NULL,
	    salary INTEGER NOT NULL
	)`

	if _, err := db.Exec(createSql); err != nil {
		return err
	}

	var count int
	if err := db.Get(&count, "SELECT COUNT(*) FROM employees"); err != nil {
		return err
	}
	if count > 0 {
		fmt.Println("已经填充测试数据")
		return nil
	}

	testData := []Employee{
		{1, "张三", "技术部", 3000},
		{2, "李四", "技术部", 2000},
		{3, "王二", "人事部", 1000},
		{4, "麻子", "技术部", 8000},
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	insertStr := "INSERT INTO employees (name,department,salary) VALUES (:name,:department,:salary)"
	if _, err := tx.NamedExec(insertStr, testData); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	fmt.Println("测试数据插入成功")

	return nil
}

// 查询技术部所有员工
func queryEmployees(db *sqlx.DB) ([]Employee, error) {
	var ems []Employee
	querySql := "SELECT id,name,department,salary From employees WHERE department =?"
	err := db.Select(&ems, querySql, "技术部")
	if err != nil {
		return nil, err
	}

	return ems, nil
}

// 查询工资最高的员工
func querySalaryMax(db *sqlx.DB) (*Employee, error) {
	var emp Employee
	querySql := "SELECT id,name,department,salary FROM employees ORDER BY salary DESC LIMIT 1"
	if err := db.Get(&emp, querySql); err != nil {
		return nil, err
	}
	return &emp, nil
}

func Sqlx1Run() {

	//dsn := "root@123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True"
	//
	//var err error
	//db, err := sqlx.Connect("mysql", dsn)
	//if err != nil {
	//	return fmt.Errorf("connect DB failed,err:%v\n", err)
	//}

	db, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		fmt.Println("数据连接失败:", err)
	}
	defer db.Close()

	if err := insertData(db); err != nil {
		fmt.Println("插入数据失败:", err)
	}

	emps, err := queryEmployees(db)
	if err != nil {
		fmt.Println("查询技术部员工失败，", err)
	}

	fmt.Println("查询技术部员工成功")
	for _, emp := range emps {
		fmt.Printf("员工ID：%d ,姓名：%s ,部门：%s ,工资：%d \n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	highest, err := querySalaryMax(db)
	if err != nil {
		fmt.Println("查询最高工资员工失败")
	}
	fmt.Printf("查询最高工资员工：%v", highest)
}

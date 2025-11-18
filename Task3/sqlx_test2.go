package Task3

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func insertBook(db *sqlx.DB) error {
	createSql := `
CREATE TABLE IF NOT EXISTS books (
    id INTEGER  PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    price REAL NOT NULL
)`
	_, err := db.Exec(createSql)
	if err != nil {
		return err
	}

	var count int
	if err := db.Get(&count, "SELECT COUNT(*) FROM books"); err != nil {
		return err
	}

	if count != 0 {
		fmt.Println("已经填充测试数据")
		return nil
	}

	testData := []Book{
		{Title: "平凡的世界", Author: "路遥", Price: 150},
		{Title: "围城", Author: "钱锺书", Price: 120},
		{Title: "月亮与六便士", Author: "毛姆", Price: 610},
		{Title: "人类简史", Author: "尤瓦尔·赫拉利", Price: 530},
		{Title: "Go语言编程", Author: "许式伟", Price: 79.00},
		{Title: "Python核心编程", Author: "Wesley Chun", Price: 128.00},
		{Title: "算法导论", Author: "Thomas Cormen", Price: 129.00},
		{Title: "设计模式", Author: "GoF", Price: 45.00},
		{Title: "重构", Author: "Martin Fowler", Price: 69.00},
		{Title: "代码大全", Author: "Steve McConnell", Price: 89.00},
		{Title: "程序员修炼之道", Author: "Andy Hunt", Price: 55.00},
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	insertSql := "INSERT INTO books (title, author, price) VALUES (:title, :author, :price)"
	_, err = db.NamedExec(insertSql, testData)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("插入测试数据成功")
	return nil

}

// 根据作者查询相关的书籍
func queryBookByAuthor(db *sqlx.DB, author string) ([]Book, error) {
	querySql := "SELECT id,title,author,price FROM books where author like ?"
	rows, err := db.Queryx(querySql, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		var book Book
		err := rows.StructScan(&book)
		if err != nil {
			continue
		}
		books = append(books, book)
	}
	return books, nil
}

func queryBookByID(db *sqlx.DB, id int) (*Book, error) {
	querySql := "SELECT id,title,author,price FROM books where id =?"
	var book Book
	if err := db.Get(&book, querySql, id); err != nil {
		return nil, err
	}
	return &book, nil
}

// 根据价格查询书籍
func queryBooksByPrice(db *sqlx.DB, price float64) ([]Book, error) {
	querySql := "SELECT id,title, author, price FROM books where price > ? ORDER BY  price DESC"

	//流式读取
	rows, err := db.Queryx(querySql, price)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book

	for rows.Next() {
		var book Book
		err = rows.StructScan(&book)
		if err != nil {
			continue
		}
		books = append(books, book)
	}

	return books, nil
}

func sqlxRun2() {
	db, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		fmt.Println("连接数据失败", err)
		return
	}
	defer db.Close()

	if err := insertBook(db); err != nil {
		fmt.Println("插入数据失败", err)
	}

	fmt.Println("插入数据成功")

	//查询单价大于50的数据
	price := 50.0
	if books, err := queryBooksByPrice(db, price); err != nil {
		fmt.Println("查询价格大约50.0的书籍失败")
	} else {
		for _, book := range books {
			fmt.Printf("查询售价大于%.2f的书有：%v\n", price, book)
		}
	}

	//查询单价大于50的数据
	author := "GOF"
	if books, err := queryBookByAuthor(db, author); err != nil {
		fmt.Println("查询GOF的书籍失败")
	} else {
		for _, book := range books {
			fmt.Printf("查询%s的书有：%v\n", author, book)
		}
	}

	//查询ID的数据
	id := 3
	if book, err := queryBookByID(db, id); err != nil {
		fmt.Printf("查询ID=%d的书籍失败", id)
	} else {
		fmt.Printf("查询id=%d的书有：%v\n", id, book)
	}
}

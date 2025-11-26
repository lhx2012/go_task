package Task3

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

//进阶gorm
//题目1：模型定义
//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//要求 ：
//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//编写Go代码，使用Gorm创建这些模型对应的数据库表。
//题目2：关联查询
//基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息。
//题目3：钩子函数
//继续使用博客系统的模型。
//要求 ：
//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

type User struct {
	ID        int `gorm:"primaryKey;autoIncrement;not null"`
	Name      string
	Posts     []Post `gorm:"foreignkey:UserID"`
	PostCount int
}

type Post struct {
	ID        int `gorm:"primaryKey;autoIncrement;not null"`
	Title     string
	Content   string
	UserID    int
	WordCount int
	State     string
	Comments  []Comment `gorm:"foreignkey:PostID"`
}

func (p *Post) BeforeCreate(db *gorm.DB) error {
	p.WordCount = len([]rune(p.Content))
	p.State = "正常"
	return nil
}

func (p *Post) AfterCreate(db *gorm.DB) error {
	var user User
	err := db.Model(&user).Where("id =?", p.UserID).First(&user).Error
	if err != nil {
		return err
	}

	user.PostCount++
	//更新文章计数
	//db.Model(&User{ID: user.ID}).Update("post_count", user.PostCount)

	if err := db.Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

type Comment struct {
	ID      int `gorm:"primaryKey;autoIncrement;not null"`
	PostID  int
	Content string
}

func (c Comment) AfterDelete(db *gorm.DB) error {
	var count int64
	err := db.Model(&Comment{}).Where("post_id =?", c.PostID).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		db.Model(&Post{}).Where("id = ?", c.PostID).Update("state", "无评论")
	}
	return nil
}

func createAndInsertDate(db *gorm.DB) {

	//创建三张表
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	var count int64
	db.Model(&User{}).Count(&count)
	if count > 0 {
		fmt.Printf("已经插入测试数据了")
	}

	//构造测试数据
	user := []User{
		{
			Name: "Alice", Posts: []Post{
				{
					Title:     "Title1",
					Content:   "Content1",
					UserID:    1,
					WordCount: 8,
					State:     "",
					Comments: []Comment{
						{PostID: 1, Content: "Comment1"},
						{PostID: 1, Content: "Comment2"},
						{PostID: 1, Content: "Comment3"},
						{PostID: 1, Content: "Comment4"},
						{PostID: 1, Content: "Comment5"},
						{PostID: 1, Content: "Comment6"},
					},
				},
				{
					Title:     "Title2",
					Content:   "Content2",
					UserID:    1,
					WordCount: 8,
					State:     "",
					Comments: []Comment{
						{PostID: 2, Content: "Comment1"},
						{PostID: 2, Content: "Comment2"},
						{PostID: 2, Content: "Comment3"},
						{PostID: 2, Content: "Comment4"},
						{PostID: 2, Content: "Comment5"},
						{PostID: 2, Content: "Comment6"},
					},
				},
			}},
		{
			Name: "Bob", Posts: []Post{
				{
					Title:     "Title3",
					Content:   "Content3",
					UserID:    2,
					WordCount: 8,
					State:     "",
					Comments: []Comment{
						{PostID: 3, Content: "Comment1"},
						{PostID: 3, Content: "Comment2"},
						{PostID: 3, Content: "Comment3"},
						{PostID: 3, Content: "Comment4"},
						{PostID: 3, Content: "Comment5"},
						{PostID: 3, Content: "Comment6"},
					},
				},
			}},
		{
			Name: "Charlie", Posts: []Post{
				{
					Title:     "Title4",
					Content:   "Content4",
					UserID:    3,
					WordCount: 8,
					State:     "",
					Comments: []Comment{
						{PostID: 4, Content: "Comment1"},
						{PostID: 4, Content: "Comment2"},
						{PostID: 4, Content: "Comment3"},
						{PostID: 4, Content: "Comment4"},
						{PostID: 4, Content: "Comment5"},
						{PostID: 4, Content: "Comment6"},
					},
				},
			}},
		{
			Name: "David", Posts: []Post{
				{
					Title:     "Title5",
					Content:   "Content5",
					UserID:    4,
					WordCount: 8,
					State:     "",
					Comments: []Comment{
						{PostID: 5, Content: "Comment1"},
						{PostID: 5, Content: "Comment2"},
						{PostID: 5, Content: "Comment3"},
						{PostID: 5, Content: "Comment4"},
						{PostID: 5, Content: "Comment5"},
						{PostID: 5, Content: "Comment6"},
					},
				},
			}},
		{
			Name: "Jack", Posts: []Post{
				{
					Title:     "Title6",
					Content:   "Content6",
					UserID:    5,
					WordCount: 8,
					State:     "",
					Comments: []Comment{
						{PostID: 6, Content: "Comment1"},
						{PostID: 6, Content: "Comment2"},
						{PostID: 6, Content: "Comment3"},
						{PostID: 6, Content: "Comment4"},
						{PostID: 6, Content: "Comment5"},
						{PostID: 6, Content: "Comment6"},
						{PostID: 6, Content: "Comment7"},
						{PostID: 6, Content: "Comment8"},
					},
				},
			}},
	}

	db.Create(&user)
}

func queryAllPostAndCommentByUser(db *gorm.DB, name string) (*User, error) {
	var user User
	err := db.Preload("Posts").Preload("Posts.Comments").Where("name =?", name).Find(&user).Error

	//err := db.Preload(clause.Associations).Find(&user, "name =?", name).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func queryMostCommentedPost(db *gorm.DB) (*Post, error) {
	var post Post
	err := db.Preload("Comments").Model(&post).Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").Group("posts.id").
		Order("comment_count DESC").First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func deleteCommentByID(db *gorm.DB, commentID int) error {
	//检查是否有这个评论
	var cmt Comment
	if err := db.Where("id = ?", commentID).First(&cmt).Error; err != nil {
		return err
	}

	if err := db.Delete(&cmt).Error; err != nil {
		return err
	}
	return nil
}

func blogRun() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//迁移表并且插入测试数据
	createAndInsertDate(db)

	//查询Alice所有的文章，跟所有的文章评论
	name := "Alice"
	user, err := queryAllPostAndCommentByUser(db, name)
	if err != nil {
		fmt.Printf("查询%s失败错误信息：%v\n", name, err)
	}

	fmt.Printf("用户%s的文章如下：\n", name)
	for _, post := range user.Posts {
		fmt.Printf("文章标题(%s),内容(%s),字数(%d)\n", post.Title, post.Content, post.WordCount)
		if post.Comments != nil {
			fmt.Println("评论：")
			for _, comment := range post.Comments {
				fmt.Printf("       %v\n", comment.Content)
			}
		}
	}

	//查询所有文章评论最多的文章
	mostPost, err := queryMostCommentedPost(db)
	if err != nil {
		fmt.Println("查找最多评论失败：", err)
	}

	fmt.Println("查找文章最多评论成功：")
	fmt.Printf("文章标题(%s),内容(%s),字数(%d)\n", mostPost.Title, mostPost.Content, mostPost.WordCount)
	if mostPost.Comments != nil {
		fmt.Println("评论：")
		for _, comment := range mostPost.Comments {
			fmt.Printf("       %v\n", comment.Content)
		}
	}

	deleteCommentByID(db, 6)
	deleteCommentByID(db, 7)
}

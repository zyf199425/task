package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Student struct {
	ID        uint   `gorm:"primary key;autoIncrement"`
	Name      string `gorm:"type:varchar(16)"`
	Age       int
	Grade     string     `gorm:"type varchar(16)"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

// SQL 练习
func sqlExercise(db *gorm.DB) {

	sqlExerTopic1(db)

	sqlExerciseTopic2(db)
}
func sqlExerTopic1(db *gorm.DB) {
	var stud Student
	// 创建表结构
	db.AutoMigrate(&stud)
	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	sql1 := "insert into students(name, age, grade) values (?, ?, ?)"
	db.Exec(sql1, "张三", 20, "三年级")

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息
	var students []Student
	sql2 := "select * from students where age > ?"
	db.Raw(sql2, 18).Scan(&students)
	fmt.Println(students)

	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	sql3 := "update students set grade = ? where  name = ?"
	db.Exec(sql3, "四年级", "张三")

	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	sql4 := "delete from students where age < ?"
	db.Exec(sql4, 15)
}

/*
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
type Account struct {
	ID        uint       `gorm:"primary key;autoIncrement"`
	Name      string     `gorm:"type:varchar(16)"`
	Balance   float64    `gorm:"type:decimal(10,2)"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}
type Transaction struct {
	ID            uint `gorm:"primary key;autoIncrement"`
	FromAccountId uint
	ToAccountId   uint
	Amount        float64    `gorm:"type:decimal(10,2)"`
	CreatedAt     *time.Time `gorm:"autoCreateTime"`
	UpdatedAt     *time.Time `gorm:"autoUpdateTime"`
}

func sqlExerciseTopic2(db *gorm.DB) {
	// 创建表结构
	db.AutoMigrate(&Account{}, &Transaction{})

	// 插入两条数据
	// accountA := Account{Name: "A", Balance: 101.0}
	// accountB := Account{Name: "B", Balance: 1.0}
	// db.Create(&accountA)
	// db.Create(&accountB)

	// 开启事务
	db.Transaction(func(tx *gorm.DB) error {
		var accountA Account
		tx.Where("name = ?", "A").First(&accountA)
		if accountA.Balance < float64(100) {
			return fmt.Errorf("余额不足")
		}
		var accountB Account
		tx.Where("name = ?", "B").First(&accountB)
		// 更新账户A的余额
		tx.Model(&accountA).Update("balance", gorm.Expr("balance - ?", 100))

		// 更新账户B的余额
		tx.Model(&accountB).Update("balance", gorm.Expr("balance + ?", 100))

		tx.Create(&Transaction{
			FromAccountId: accountA.ID,
			ToAccountId:   accountB.ID,
			Amount:        100,
		})
		return nil
	})
}

type Employee struct {
	ID         uint       `gorm:"primary key ; autoIncrement"`
	Name       string     `gorm:"type:varchar(16)"`
	Department string     `gorm:"type:varchar(32)"`
	Salary     float64    `gorm:"type:decimal(10,2)"`
	CreatedAt  *time.Time `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time `gorm:"autoUpdateTime"`
}

func sqlxExercise(db *gorm.DB) {
	db.AutoMigrate(&Employee{})

	// 插入数据
	employees := []Employee{{Name: "张三", Department: "研发部", Salary: 10000}, {Name: "李四", Department: "人力资源部", Salary: 6000}, {Name: "王五", Department: "技术部", Salary: 9000}}
	db.Create(&employees)

	// 查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	var employeeList []Employee
	db.Where("department = ?", "技术部").Find(&employeeList)
	fmt.Println(employeeList)

	// 查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var employee Employee
	db.Order("salary desc").First(&employee)
	fmt.Println(employee)
}

type Book struct {
	ID        uint       `gorm: "primary key; autoIncrement"`
	Title     string     `gorm:"type:varchar(16)"`
	Author    string     `gorm:"type:varchar(16)"`
	Price     float64    `gorm:"type:decimal(10,2)"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func sqlxExercise2(db *gorm.DB) {
	db.AutoMigrate(&Book{})
	books := []Book{{Title: "西游记", Author: "吴承恩", Price: 20.0}, {Title: "红楼梦", Author: "曹雪芹", Price: 30.0}, {Title: "水浒传", Author: "施耐庵", Price: 25.0}}
	db.Create(&books)

	var bookList []Book
	db.Where("price > ? and price < ?", 20.0, 30.0).Find(&bookList)
	fmt.Println(bookList)

}

// 进阶gorm
type User struct {
	ID        uint   `gorm:"primary key;autoIncrement"`
	Name      string `gorm:"type:varchar(16)"`
	Age       int
	Sex       int `gorm:"type:tinyint"`
	PostCount int
	Posts     []Post
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}
type Post struct {
	ID            uint   `gorm:"primary key;autoIncrement"`
	Title         string `gorm:"type:varchar(16)"`
	Content       string `gorm:"type:text"`
	UserID        uint
	User          User
	CommentStatus int `gorm:"type:tinyint"`
	CommentCount  int
	Comments      []Comment
	CreatedAt     *time.Time `gorm:"autoCreateTime"`
	UpdatedAt     *time.Time `gorm:"autoUpdateTime"`
}
type Comment struct {
	ID        uint   `gorm:"primary key;autoIncrement"`
	Content   string `gorm:"type:text"`
	PostID    uint
	Post      Post
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

// 钩子函数
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	// 更新用户的文章数量
	tx.Model(&User{}).Where("ID = ?", p.UserID).Update("PostCount", gorm.Expr("post_count + ?", 1))
	return
}
func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	// 更新文章的评论数量和评论状态
	tx.Model(&Post{}).Select("CommentStatus", "CommentCount").Where("ID = ?", c.PostID).Updates(map[string]interface{}{"CommentStatus": 1, "CommentCount": gorm.Expr("comment_count + ?", 1)})
	return
}
func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	// 更新文章的评论数量和评论状态
	tx.Model(&Post{}).Where("ID = ?", c.PostID).Update("CommentCount", gorm.Expr("comment_count - ?", 1))
	var post Post
	tx.First(&post, c.PostID)
	if post.CommentCount == 0 {
		tx.Model(&post).Select("CommentStatus").Update("comment_status", 0)
	}
	return
}

func gormExercise(db *gorm.DB) {
	// 1 创建表结构
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// 2 插入数据
	users := []User{{Name: "张三", Age: 18, Sex: 1, PostCount: 0}, {Name: "李四", Age: 20, Sex: 0, PostCount: 0}, {Name: "王五", Age: 22, Sex: 1, PostCount: 0}}
	db.Create(&users)
	posts := []Post{{Title: "第一篇博客", Content: "这是一篇测试博客", UserID: 1, CommentStatus: 0, CommentCount: 0}, {Title: "第二篇博客", Content: "这是第二篇测试博客", UserID: 2, CommentStatus: 0, CommentCount: 0}, {Title: "第三篇博客", Content: "这是第三篇测试博客", UserID: 3, CommentStatus: 0, CommentCount: 0}}
	db.Create(&posts)
	comments := []Comment{{Content: "这是第条7评论", PostID: 7}}
	db.Create(&comments)

	// 查询某个用户发布的所有文章及其对应的评论信息。
	var user User
	db.Preload("Posts").Preload("Posts.Comments").Where("id = ?", 1).First(&user)
	fmt.Println(user)

	// // 查询评论数最多的文章
	var post Post
	db.Preload("Comments").Order("comment_count desc").First(&post)
	fmt.Println("评论数最高的文章：", post)

	// 测试删除钩子
	db.Delete(&Comment{ID: 14, PostID: 8})
}

func main() {
	dsn := "root:ZhaoYunFei@0205@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	fmt.Println(db, err)
	if err != nil {
		panic("连接数据库失败！")
	}

	sqlExercise(db)
	sqlxExercise(db)
	sqlxExercise2(db)
	gormExercise(db)
}

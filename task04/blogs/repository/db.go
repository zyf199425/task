package repository

import (
	"errors"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID        uint       `gorm:"primary key;autoIncrement" json:"id"`
	UserName  string     `gorm:"type:varchar(16)" json:"userName"`
	Password  string     `gorm:"type:varchar(32)" json:"password"`
	Email     string     `gorm:"type:varchar(32)" json:"email"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type Post struct {
	ID        uint       `gorm:"primary key;autoIncrement" json:"id"`
	Title     string     `gorm:"type:varchar(16)" json:"title"`
	Content   string     `gorm:"type:text" json:"content"`
	UserId    uint       `gorm:"type:uint" json:"userId"`
	User      User       `json:"user"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type Comment struct {
	ID        uint       `gormm:"primary key;autoIncrement" json:"id"`
	Content   string     `gorm:"type:text" json:"content"`
	UserId    uint       `gorm:"type:uint" json:"userId"`
	User      User       `json:"user"`
	PostId    uint       `gorm:"type:uint" json:"postId"`
	Post      Post       `json:"post"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

var DB *gorm.DB

func init() {
	// 创建表结构
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
	if err != nil {
		panic("连接数据库失败！")
	}
	DB = db
	// 创建表结构
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

func QueryUserByName(username string) User {
	var u User
	DB.Model(&User{}).Where("user_name = ?", username).First(&u)
	if u.ID == 0 {
		return User{}
	}

	return u
}

func SaveUser(user *User) (bool, error) {
	db := DB.Create(user)
	if db.RowsAffected == 0 {
		return false, errors.New("保存失败！")
	}
	return true, nil
}

func SavePost(post *Post) (bool, error) {
	db := DB.Create(post)
	if db.RowsAffected == 0 {
		return false, errors.New("保存失败！")
	}
	return true, nil
}

func UpdatePostById(post *Post) bool {
	db := DB.Model(&Post{}).Where("id = ?", post.ID).Updates(post)
	if db.RowsAffected == 0 {
		return false
	}
	return true
}

func DeletePostById(id uint) bool {
	db := DB.Delete(&Post{}, id)
	if db.RowsAffected == 0 {
		return false
	}
	return true
}

func QueryPostById(id uint) Post {
	var post Post
	DB.Preload("User").Model(&Post{}).Where("id = ?", id).First(&post)
	return post
}

func QueryPostPage(pageNum int, pageSize int) []Post {
	offset := (pageNum - 1) * pageSize
	var posts []Post
	DB.Offset(offset).Limit(pageSize).Find(&posts)
	return posts
}

func SaveComment(comment *Comment) (bool, error) {
	db := DB.Create(comment)
	if db.RowsAffected == 0 {
		return false, errors.New("保存失败！")
	}
	return true, nil
}

func QueryCommentByPostId(postId uint) []Comment {
	var comments []Comment
	DB.Preload("User").Preload("Post").Preload("Post.User").Model(&Comment{}).Where("post_id = ?", postId).Find(&comments)
	return comments
}

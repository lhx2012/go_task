package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Role     string    `gorm:"size:32;default:user"`
	Posts    []Post    `gorm:"foreignkey:user_id"`
	Comments []Comment `gorm:"foreignkey:user_id"`
}

type UserLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResp struct {
	Token string `json:"token"`
}

type UserPageReq struct {
	Page     int `form:"page" binding:"min=1"`
	PageSize int `form:"pageSize" binding:"min=1"`
}

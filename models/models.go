package models

import "gorm.io/gorm"

// User — jo log blog likhte hain
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"` // response mein nahi aayega
}

// Post — blog post
type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"author" gorm:"foreignKey:UserID"`
}

// Register ke liye input
type RegisterInput struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login ke liye input
type LoginInput struct {
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Post banane ke liye input
type PostInput struct {
	Title   string `json:"title"   binding:"required"`
	Content string `json:"content" binding:"required"`
}
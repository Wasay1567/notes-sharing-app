package models

import (
	"time"
)

type Notes struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `gorm:"index"`
	Title      string    `json:"title" gorm:"not null"`
	Content    string    `json:"content" gorm:"not null"`
	Expiration string    `json:"expires_in" gorm:"not null"`
	Created_at time.Time `json:"created_at"`
	Viewed     bool      `json:"viewed" gorm:"default:false"`
	IsExpired  bool      `json:"is_expired" gorm:"default:false"`
	Password   string    `json:"password"`
}

type GetNote struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Expiration string `json:"expires_in"`
	Password   string `json:"password"`
}

type Resp struct {
	Id  uint   `json:"id"`
	Url string `json:"url"`
}

type ReadNote struct {
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	UserNotes []Notes `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

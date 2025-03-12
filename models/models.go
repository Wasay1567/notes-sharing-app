package models

type Notes struct {
	Id         string `json:"id" gorm:"primaryKey"`
	Title      string `json:"title" gorm:"not null"`
	Content    string `json:"content" gorm:"not null"`
	Expiration string `json:"expires_in" gorm:"not null"`
	Created_at string `json:"created_at"`
	Password   string `json:"password" gorm:"not null"`
}

type GetNote struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Expiration string `json:"expires_in"`
	Password   string `json:"password"`
}

type Resp struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

var NotesList = []Notes{}

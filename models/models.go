package models



type Notes struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Expiration string `json:"expires_in"`
	Created_at string `json:"created_at"`
	Password   string `json:"password"`
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

package model

import "gorm.io/gorm"

/* assets的资源文件：https://getbootstrap.com/ https://pandao.github.io/editor.md/ */
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Post struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text"`
	Tag     string
}

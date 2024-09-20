package models

import (
	"gorm.io/gorm"
)

type PostType string

const (
	Image PostType = "image"
	Link  PostType = "link"
	Video PostType = "video"
	File  PostType = "file"
)

type Visibility string

const (
	Published Visibility = "published"
	Private   Visibility = "private"
)

type Post struct {
	gorm.Model
	Type          PostType `gorm:"type:varchar(20)"`
	Description   string
	Thoughts      []Thought
	PublishStatus string     `gorm:"type:varchar(20)"`
	Visibility    Visibility `gorm:"type:varchar(20)"`
	Tags          []Tag      `gorm:"many2many:post_tags;"`
}

type Thought struct {
	gorm.Model
	PostID  uint
	Content string
}

type Tag struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex"`
	Posts []Post `gorm:"many2many:post_tags;"`
}

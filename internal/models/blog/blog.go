package modelblog

import "backend/internal/config"

type Content struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	Title     string
	AuthorId  int
	CreatedAt string
	EditAt    string
	ViewCount string
	Content   string
	Tags      string
	State     int
}

func (Content) TableName() string {
	return config.TABLENAME_BLOG["Content"]
}

type UserBlogView struct {
	UserId    int
	BlogId    int
	CreatedAt string
}

func (UserBlogView) TableName() string {
	return config.TABLENAME_BLOG["UserBlogView"]
}

type UserBlogVoteUp struct {
	UserId    int
	BlogId    int
	CreatedAt string
}

func (UserBlogVoteUp) TableName() string {
	return config.TABLENAME_BLOG["UserBlogVoteUp"]
}

type UserBlogVoteDown struct {
	UserId    int
	BlogId    int
	CreatedAt string
}

func (UserBlogVoteDown) TableName() string {
	return config.TABLENAME_BLOG["UserBlogVoteDown"]
}

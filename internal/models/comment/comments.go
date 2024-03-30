package modelcomment

import "backend/internal/config"

type Comment struct {
	Id            int `gorm:"primaryKey;autoIncrement"`
	ReferId       int
	BlogId        int
	UserId        int
	UserName      string
	Identity      string
	Content       string
	CreatedAt     string
	VoteUpCount   int
	VoteDownCount int
	State         int
}

func (Comment) TableName() string {
	return config.TABLENAME_COMMENT["Comment"]
}

type UserCommentVoteUp struct {
	UserId    int
	CommentId int
	CreatedAt string
}

func (UserCommentVoteUp) TableName() string {
	return config.TABLENAME_COMMENT["UserCommentVoteUp"]
}

type UserCommentVoteDown struct {
	UserId    int
	CommentId int
	CreatedAt string
}

func (UserCommentVoteDown) TableName() string {
	return config.TABLENAME_BLOG["UserCommentVoteDown"]
}

package repoblog

import (
	modelblog "backend/internal/models/blog"
	"backend/internal/repositories"
	utillog "backend/internal/util/log"
	"log"
)

const blogPath = "repositories/blog"

type userIdStruct struct {
	UserId    int
	CreatedAt string
}
type blogIdStruct struct {
	BlogId    int
	CreatedAt string
}

func GetBlog(id int) modelblog.Content {
	const funcName = "GetBlog"

	blogContent := modelblog.Content{}
	err := repositories.BlogDB.Model(&modelblog.Content{}).Where("id = ?", id).First(&blogContent).Error
	if err != nil {
		utillog.FormatString(blogPath, funcName, "查询数据库失败，此文章不存在")
		log.Println("WARN: " + err.Error())
	}
	return blogContent
}

func GetVoteInfo(blogId int, voteType string) ([]userIdStruct, error) {
	const funcName = "GetBlogVoteUp"

	var userIdSlice []userIdStruct
	var err error

	switch voteType {
	case "up":
		err = repositories.BlogDB.Model(&modelblog.UserBlogVoteUp{}).Where("blog_id = ?", blogId).Scan(&userIdSlice).Error

	case "down":
		err = repositories.BlogDB.Model(&modelblog.UserBlogVoteDown{}).Where("blog_id = ?", blogId).Scan(&userIdSlice).Error
	}

	if err != nil {
		utillog.FormatString(blogPath, funcName, "查询点赞列表时出现错误")
		log.Println("WARN: " + err.Error())
	}
	return userIdSlice, err
}

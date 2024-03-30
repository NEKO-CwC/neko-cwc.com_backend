package repoblog

import (
	modelblog "backend/internal/models/blog"
	"backend/internal/repositories"
	utillog "backend/internal/util/log"
	"log"
	"time"
)

func Vote(userId int, blogId int, voteType string, operateType string) error {
	var funcName = "Vote"
	var err error
	var model any

	voteInfo := modelblog.UserBlogVoteUp{
		UserId:    userId,
		BlogId:    blogId,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// 确定使用哪个模型
	switch voteType {
	case "up":
		model = modelblog.UserBlogVoteUp{}
	case "down":
		model = modelblog.UserBlogVoteDown{}
	}

	// 执行对应操作
	switch operateType {
	case "cancel":
		err = repositories.BlogDB.Model(model).Delete(voteInfo).Error
	case "confirm":
		err = repositories.BlogDB.Model(model).Create(voteInfo).Error
	}

	if err != nil {
		utillog.FormatString(blogPath, funcName, "处理记录出错")
		log.Println("WARN:" + err.Error())
	}
	return err
}

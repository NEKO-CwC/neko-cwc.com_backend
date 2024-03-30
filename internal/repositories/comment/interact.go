package repocomment

import (
	modelcomment "backend/internal/models/comment"
	modeluser "backend/internal/models/user"
	"backend/internal/repositories"
	utillog "backend/internal/util/log"
	"log"
	"time"
)

type voteInfo struct {
	UserId    int
	CommentId int
	CreatedAt string
}

func Vote(userId int, commentId int, voteType string, operateType string) error {
	var funcName = "Vote"

	var err error
	var model any

	info := voteInfo{
		UserId:    userId,
		CommentId: commentId,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	switch voteType {
	case "up":
		model = modelcomment.UserCommentVoteUp{}
	case "down":
		model = modelcomment.UserCommentVoteDown{}
	}

	switch operateType {
	case "confirm":
		err = repositories.CommentDB.Model(&model).Create(&info).Error
	case "cancel":
		err = repositories.CommentDB.Model(&model).Delete(&info).Error
	}
	if err != nil {
		utillog.FormatString(commentPath, funcName, "处理点赞信息时出现问题")
		log.Println("WARN: " + err.Error())
	}
	return err
}

func Reply(userInfo modeluser.Info, blogId int, referId int, content string) error {
	var funcName = "Reply"

	var err error
	commentInfo := modelcomment.Comment{
		ReferId:       referId,
		BlogId:        blogId,
		UserId:        userInfo.Id,
		UserName:      userInfo.Name,
		Identity:      userInfo.Identity,
		Content:       content,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		VoteUpCount:   0,
		VoteDownCount: 0,
		State:         1,
	}
	err = repositories.CommentDB.Model(&modelcomment.Comment{}).Create(&commentInfo).Error
	if err != nil {
		utillog.FormatString(commentPath, funcName, "创建新评论出错")
		log.Println("WARN: " + err.Error())
	}
	return err
}

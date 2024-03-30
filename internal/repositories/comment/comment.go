package repocomment

import (
	modelcomment "backend/internal/models/comment"
	"backend/internal/repositories"
	utillog "backend/internal/util/log"
	"log"
)

const commentPath = "repositories/comment"

type VoteInfo struct {
	UserId    int
	CreatedAt string
}

func GetComments(blogId int) ([]modelcomment.Comment, error) {
	const funcName = "GetComments"

	var commentSlice []modelcomment.Comment
	err := repositories.CommentDB.Model(&modelcomment.Comment{}).Where("blog_id = ?", blogId).Find(&commentSlice).Error
	if err != nil {
		utillog.FormatString(commentPath, funcName, "查询评论时出现错误")
		log.Println("WARN: " + err.Error())
	}
	return commentSlice, err
}
func GetVoteInfo(userId int, voteType string) ([]VoteInfo, error) {
	const funcName = "GetUserVoteUp"

	var blogIdSlice []VoteInfo
	var err error

	switch voteType {
	case "up":
		err = repositories.CommentDB.Model(&modelcomment.UserCommentVoteUp{}).Where("user_id = ?", userId).Scan(&blogIdSlice).Error
	case "down":
		err = repositories.CommentDB.Model(&modelcomment.UserCommentVoteDown{}).Where("user_id = ?", userId).Scan(&blogIdSlice).Error
	}
	if err != nil {
		utillog.FormatString(commentPath, funcName, "查询点赞列表时出现错误")
		log.Println("WARN:" + err.Error())
	}
	return blogIdSlice, err
}

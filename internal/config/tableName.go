package config

var TABLENAME_BLOG = map[string]string{
	"Content":          "content",
	"UserBlogView":     "user_blog_view",
	"UserBlogVoteUp":   "user_blog_vote_up",
	"UserBlogVoteDown": "user_blog_vote_down",
}

var TABLENAME_ACCOUNT = map[string]string{
	"Info": "basic_info",
}

var TABLENAME_COMMENT = map[string]string{
	"Comment":             "comment",
	"UserCommentVoteUp":   "user_comment_vote_up",
	"UserCommentVoteDown": "user_comment_vote_down",
}

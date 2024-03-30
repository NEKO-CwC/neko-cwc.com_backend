package modeluser

import "backend/internal/config"

type Info struct {
	Id           int `gorm:"primaryKey;autoIncrement"`
	Name         string
	Email        string
	Password     string
	Salt         string
	LastLogin    string
	CreatedAt    string
	Identity     string
	Tags         string
	ArticleCount int
	State        int
}

func (Info) TableName() string {
	return config.TABLENAME_ACCOUNT["Info"]
}

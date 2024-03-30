package repouserinfo

import (
	middleware_repository "backend/internal/middleware/repository"
	"backend/internal/models/user"
	"backend/internal/repositories"
	utillog "backend/internal/util/log"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

const accountPath string = "repositories/account"

// CheckAccountExist
// 返回值有以下可能
// false：不存在账号信息
// true： 存在
func CheckAccountExist(email string) (bool, error) {
	funcName := "CheckAccountExist"

	err := repositories.UserInfoDB.Model(&modeluser.Info{}).Where("email = ?", email).First(&modeluser.Info{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("INFO: modeluser %v not exist", email)
		return false, err
	} else if err != nil {
		utillog.FormatString(accountPath, funcName, "对email进行查询")
		log.Println("WARN: " + err.Error())
		return false, err
	}
	return true, err
}

func Signup(email string, password string) (bool, error) {
	funcName := "CreatedAt"

	var err error
	var userInfo modeluser.Info

	var totalUserCount int
	var totalUserCount64 int64
	repositories.UserInfoDB.Model(&modeluser.Info{}).Count(&totalUserCount64)
	totalUserCount = int(totalUserCount64)
	userInfo.Id = totalUserCount + 1

	userInfo.Name = email
	userInfo.Email = email
	userInfo.Salt, err = middleware_repository.GenerateSalt(16)
	if err != nil {
		utillog.FormatString(accountPath, funcName, "对salt进行生成")
		log.Println("WARN: " + err.Error())
		return false, err
	}
	userInfo.Password, err = middleware_repository.HashPassword(password, userInfo.Salt)
	if err != nil {
		utillog.FormatString(accountPath, funcName, "对password进行哈希加密")
		log.Println("WARN: " + err.Error())
		return false, err
	}
	userInfo.LastLogin = time.Now().Format("2006-01-02 15:04:05")
	userInfo.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	userInfo.Identity = "user"
	userInfo.State = 1
	err = repositories.UserInfoDB.Model(&modeluser.Info{}).Create(&userInfo).Error
	if err != nil {
		utillog.FormatString(accountPath, funcName, "对用户信息进行创建")
		log.Println("WARN: " + err.Error())
		return false, err
	}
	return true, err
}

func Login(email string, password string) (modeluser.Info, error, string) {
	funcName := "Login"

	var err error
	var userInfo modeluser.Info

	exist, err := CheckAccountExist(email)
	if !exist {
		return userInfo, err, "账号不存在"
	}

	err = repositories.UserInfoDB.Model(&modeluser.Info{}).Where("email = ?", email).First(&userInfo).Error
	if err != nil {
		utillog.FormatString(accountPath, funcName, "对email进行查找")
		log.Println("WARN: " + err.Error())
		return userInfo, err, "数据库错误"
	}
	localHashedPassword := userInfo.Password
	salt := userInfo.Salt
	inputHashedPassword, err := middleware_repository.HashPassword(password, salt)
	if err != nil {
		utillog.FormatString(accountPath, funcName, "对inputPassword进行哈希加密")
		log.Println("WARN: " + err.Error())
		return userInfo, err, "后端行为错误"
	}
	if inputHashedPassword == localHashedPassword {
		userInfo.LastLogin = time.Now().Format("2006-01-02 15:04:05")
		err = repositories.UserInfoDB.Model(&modeluser.Info{}).Where("id = ?", userInfo.Id).Updates(&userInfo).Error
		if err != nil {
			utillog.FormatString(accountPath, funcName, "对userInfo中的数据进行写入")
			log.Println("WARN: " + err.Error())
			return userInfo, err, "数据库错误"
		}
	} else {
		return userInfo, err, "密码错误"
	}
	return userInfo, err, ""
}

func CheckUserInfo(id int) (modeluser.Info, error) {
	funcName := "CheckUserInfo"
	var err error
	var userInfo modeluser.Info

	err = repositories.UserInfoDB.Model(&modeluser.Info{}).Where("id = ?", id).First(&userInfo).Error
	if err != nil {
		utillog.FormatString(accountPath, funcName, "连接数据库出现错误")
		log.Println("WARN: " + err.Error())
		return userInfo, err
	}
	return userInfo, nil
}

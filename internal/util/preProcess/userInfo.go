package preProcess

import modeluser "backend/internal/models/user"

func BasicInfo(info modeluser.Info) modeluser.Info {
	info.Password = ""
	info.Salt = ""
	return info
}

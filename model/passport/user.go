package passport

import (
	"errors"

	"tdp-cloud/model/user"
	"tdp-cloud/module/midware"
	"tdp-cloud/module/upload"
)

// 登录账号

type LoginParam struct {
	Username  string `binding:"required"`
	Password  string `binding:"required"`
	IpAddress string
	UserAgent string
}

type LoginResult struct {
	UserId   uint
	Username string
	Avatar   string
	Level    uint
	AppId    string
	Email    string
	Token    string
}

func Login(data *LoginParam) (*LoginResult, error) {

	ur, _ := user.Fetch(&user.FetchParam{
		Username: data.Username,
	})

	// 验证账号

	if ur.Id == 0 {
		return nil, errors.New("账号错误")
	}
	if !user.CheckSecret(ur.Password, data.Password) {
		return nil, errors.New("密码错误")
	}

	// 创建令牌

	token, err := midware.CreateToken(&midware.UserInfo{
		Id:     ur.Id,
		Level:  ur.Level,
		AppKey: ur.AppKey,
	})

	if err != nil {
		return nil, err
	}

	// 返回结果

	res := &LoginResult{
		UserId:   ur.Id,
		Username: ur.Username,
		Avatar:   ur.Avatar,
		Level:    ur.Level,
		AppId:    ur.AppId,
		Email:    ur.Email,
		Token:    token,
	}

	return res, nil

}

// 修改资料

type ProfileUpdateParam struct {
	user.UpdateParam
	OldPassword string `binding:"required"`
}

func ProfileUpdate(data *ProfileUpdateParam) error {

	ur, _ := user.Fetch(&user.FetchParam{Id: data.Id})

	// 验证账号

	if ur.Id == 0 {
		return errors.New("账号错误")
	}
	if !user.CheckSecret(ur.Password, data.OldPassword) {
		return errors.New("密码错误")
	}
	if err := user.CheckUserinfo(data.Username, data.Password, data.Email); err != nil {
		return err
	}

	// 更新信息

	return user.Update(&data.UpdateParam)

}

// 更新头像

type AvatarUpdateParam struct {
	UserId      uint
	Base64Image string `binding:"required"`
}

func AvatarUpdate(rq *AvatarUpdateParam) (string, error) {

	filePath := upload.BaseDir + "/avatar" + upload.UintPathname(rq.UserId) + ".png"

	if err := upload.SaveBase64Image(filePath, rq.Base64Image); err != nil {
		return "", err
	}

	uu := &user.UpdateParam{
		Id:     rq.UserId,
		Avatar: filePath,
	}
	if err := user.Update(uu); err != nil {
		return "", err
	}

	return uu.Avatar, nil

}

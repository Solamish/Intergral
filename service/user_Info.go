package service

import (
	"mobileSign/model"
)

type User struct {
	ID                int    `json:"id"`
	RedId             string `json:"redid"`
	Stunum            string `json:"stunum"`
	UserName          string `json:"username"`
	NickName          string `json:"nickname"`
	CheckInDays       int    `json:"check_in_days"`
	HeadImgUrl        string `json:"photo_src"`
	PhotoThumbnailSrc string `json:"photo_thumbnail_src"` // 缩略图
	Integral          int    `json:"integral"`
	Phone             string `json:"phone"`
	Gendor            string `json:"gendor"`
	Introduction      string `json:"introduction"`
	QQ                string `json:"qq"`
}

func UserInfo(redId string) (User){
	user := model.GetUserInfo(redId)
	userInfo := User{
		ID:                user.ID,
		RedId:             user.RedId,
		NickName:          user.NickName,
		HeadImgUrl:        user.HeadImgUrl,
		CheckInDays:       user.CheckInDays,
		Integral:          user.Integral,
		Phone:             user.Phone,
		Gendor:            user.Gendor,
		Stunum:            user.Stunum,
		Introduction:      user.Introduction,
		QQ:                user.QQ,
		UserName:          user.UserName,
		PhotoThumbnailSrc: user.PhotoThumbnailSrc,
	}
	return userInfo
}

package model

import "time"

type UserModel struct {
	Id        int32     `gorm:"column:id" json:"id"`
  OpenId    string   `gorm:"column:openId" json:"openId"`
	NickName  string   `gorm:"column:nickName" json:"nickName"`
	AvatarFileId string `gorm:"column:avatarFileId" json:"avatarFileId"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}
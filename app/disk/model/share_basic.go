package model

import (
	"time"

	"gorm.io/gorm"
)

type ShareBasic struct {
	Id                     int
	Identity               string
	UserIdentity           string
	UserRepositoryIdentity string
	RepositoryIdentity     string
	ExpiredTime            int
	ClickNum               int
	Desc                   string
	CreatedAt              time.Time      `gorm:"created"`
	UpdatedAt              time.Time      `gorm:"updated"`
	DeletedAt              gorm.DeletedAt `gorm:"deleted"`
	DelState               int
}

func (ShareBasic) TableName() string {
	return "share_basic"
}

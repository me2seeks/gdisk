package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	Id                 int64
	Identity           string
	Uid                string
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
	CreatedAt          time.Time      `gorm:"created"`
	UpdatedAt          time.Time      `gorm:"updated"`
	DeletedAt          gorm.DeletedAt `gorm:"deleted"`
}

func (table UserRepository) TableName() string {
	return "user_repository"
}

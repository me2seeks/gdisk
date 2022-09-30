package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	Id           int
	Identity     string
	Uid          string
	Pid          int64
	RepositoryId string
	Ext          string
	Name         string
	CreatedAt    time.Time      `gorm:"created"`
	UpdatedAt    time.Time      `gorm:"updated"`
	DeletedAt    gorm.DeletedAt `gorm:"deleted"`
	DelState     int
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
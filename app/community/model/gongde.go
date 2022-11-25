package model

import (
	"time"

	"gorm.io/gorm"
)

type GongdeBasic struct {
	Id        int
	Count     int
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at"`
	DelState  int
}

func (table GongdeBasic) TableName() string {
	return "gongde_basic"
}

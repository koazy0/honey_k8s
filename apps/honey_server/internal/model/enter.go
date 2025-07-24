package model

import (
	"gorm.io/gorm"
	"honey_server/internal/common"
	"time"
)

var logger = common.Logs().Cat("model")

type Model struct {
	ID        uint           `gorm:"primaryKey;comment:id" json:"id,select($any)" structs:"-"` // 主键ID
	CreatedAt time.Time      `gorm:"comment:创建时间" json:"created_at,select($any)" structs:"-"`  // 创建时间
	UpdatedAt time.Time      `gorm:"comment:更新时间" json:"-" structs:"-"`                        // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Key   string `form:"key"`
}

type IDListRequest struct {
	IdList []uint `json:"idList"`
}

type IDRequest struct {
	Id uint `json:"id" uri:"id" form:"id"`
}

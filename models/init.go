package models

import (
	"github.com/aiyijing/smart-gateway/database"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func init() {
	err := database.GetInstance().AutoMigrate(
		&Node{},
	)
	if err != nil {
		panic(err)
	}
}

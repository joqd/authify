package model

import (
	"time"
)

type UserActivityModel struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	UserID    uint64    `gorm:"not null;index"`
	Event     string    `gorm:"type:text"`
	CreatedAt time.Time `goem:"autoCreateTime"`
}

func (UserActivityModel) TableName() string {
	return "user_activities"
}

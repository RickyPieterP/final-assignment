package mysql

import (
	"mygram/app/shared/constant"
	"time"
)

type Comment struct {
	Id           int       `gorm:"column:id;primary_key"`
	UserId       int       `gorm:"column:user_id"`
	PhotoId      int       `gorm:"column:photo_id"`
	Title        string    `gorm:"column:title"`
	Message      string    `gorm:"column:message"`
	Created_Date time.Time `gorm:"column:created_date"`
	Updated_At   time.Time `gorm:"column:updated_at"`
}

func (Comment) TableName() string {
	return constant.COMMENT
}

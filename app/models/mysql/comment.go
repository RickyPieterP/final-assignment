package mysql

import (
	"mygram/app/shared/constant"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	Id           int       `gorm:"column:id;primary_key"`
	UserId       int       `gorm:"column:user_id"`
	PhotoId      int       `gorm:"column:photo_id"`
	Message      string    `gorm:"column:message;not null"`
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At   time.Time `gorm:"column:updated_at"`
}

func (Comment) TableName() string {
	return constant.COMMENT
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	c.Created_At = time.Now()
	c.Updated_At = time.Now()
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	c.Updated_At = time.Now()
	return
}

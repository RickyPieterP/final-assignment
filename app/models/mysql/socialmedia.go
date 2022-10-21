package mysql

import (
	"mygram/app/shared/constant"
	"time"
)

type SocialMedia struct {
	Id             int       `gorm:"column:id;primary_key"`
	UserID         int       `gorm:"column:user_id"`
	Name           string    `gorm:"column:name"`
	SocialMediaUrl string    `gorm:"column:social_media_url"`
	Title          string    `gorm:"column:title"`
	Created_Date   time.Time `gorm:"column:created_date"`
	Updated_At     time.Time `gorm:"column:updated_at"`
}

func (SocialMedia) TableName() string {
	return constant.SOCIALMEDIA
}

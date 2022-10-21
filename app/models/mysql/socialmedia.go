package mysql

import (
	"mygram/app/shared/constant"
	"time"

	"gorm.io/gorm"
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

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	s.Created_Date = time.Now()
	s.Updated_At = time.Now()
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	s.Updated_At = time.Now()
	return
}

func (s *SocialMedia) TableName() string {
	return constant.SOCIALMEDIA
}

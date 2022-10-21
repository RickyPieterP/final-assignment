package mysql

import (
	"mygram/app/shared/constant"
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	Id             int       `gorm:"column:id;primary_key"`
	UserID         int       `gorm:"column:user_id"`
	Name           string    `gorm:"column:name;not null"`
	SocialMediaUrl string    `gorm:"column:social_media_url;not null"`
	Created_At   time.Time `gorm:"column:created_at"`
	Updated_At     time.Time `gorm:"column:updated_at"`
}

func (SocialMedia) TableName() string {
	return constant.SOCIALMEDIA
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	p.Created_At = time.Now()
	p.Updated_At = time.Now()
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	p.Updated_At = time.Now()
	return
}
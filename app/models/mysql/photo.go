package mysql

import (
	"mygram/app/shared/constant"
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	Id           int    `gorm:"column:id;primary_key"`
	UserId       int    `gorm:"column:user_id"`
	Title        string `gorm:"column:title;not null"`
	Caption      string `gorm:"column:caption;not null"`
	PhotoUrl     string `gorm:"column:photo_url;not null"`
	User         User
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At   time.Time `gorm:"column:updated_at"`
}

func (Photo) TableName() string {
	return constant.PHOTO
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	p.Created_At = time.Now()
	p.Updated_At = time.Now()
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	p.Updated_At = time.Now()
	return
}
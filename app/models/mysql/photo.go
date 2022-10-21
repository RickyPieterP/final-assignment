package mysql

import (
	"mygram/app/shared/constant"
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	Id           int    `gorm:"column:id;primary_key"`
	UserId       int    `gorm:"column:user_id"`
	Title        string `gorm:"column:title"`
	Caption      string `gorm:"column:caption"`
	PhotoUrl     string `gorm:"column:photo_url"`
	User         User
	Created_Date time.Time `gorm:"column:created_date"`
	Updated_At   time.Time `gorm:"column:updated_at"`
}

func (Photo) TableName() string {
	return constant.PHOTO
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	p.Created_Date = time.Now()
	p.Updated_At = time.Now()
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	p.Updated_At = time.Now()
	return
}
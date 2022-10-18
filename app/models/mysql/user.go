package mysql

import (
	"mygram/app/shared/constant"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           int       `gorm:"column:id;primary_key"`
	Username     string    `gorm:"column:username"`
	Email        string    `gorm:"column:email"`
	Password     string    `gorm:"column:password"`
	Age          int       `gorm:"column:age"`
	Created_Date time.Time `gorm:"column:created_date"`
	Updated_At   time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return constant.USER
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Created_Date = time.Now()
	u.Updated_At = time.Now()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.Updated_At = time.Now()
	return
}
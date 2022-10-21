package mysql

import (
	"mygram/app/shared/constant"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           int       `gorm:"column:id;primary_key"`
	Username     string    `gorm:"column:username;uniqueIndex"`
	Email        string    `gorm:"column:email;uniqueIndex;not null"`
	Password     string    `gorm:"column:password;not null"`
	Age          int       `gorm:"column:age;not null"`
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At   time.Time `gorm:"column:updated_at"`
}


func (User) TableName() string {
	return constant.USER
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Created_At = time.Now()
	u.Updated_At = time.Now()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.Updated_At = time.Now()
	return
}


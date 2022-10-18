package database

import (
	"fmt"
	"mygram/app/shared/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// models "mygram/app/models/mysql"
)

type MySQL struct {
	*gorm.DB
}

func SetupMySQL() *MySQL {
	fmt.Println("Try Setup MySQL ...")

	url := config.DB.MySQL.DataSourceName
	// dialect := config.DB.MySQL.DriverName

	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		panic(err)
	}
	// Enable Logger, show detailed log
	db.Logger.LogMode(4)

	// db.AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.SocialMedia{})

	// Connection Pool
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(40)
	sqlDB.SetConnMaxIdleTime(50 * time.Minute)
	sqlDB.SetConnMaxLifetime(50 * time.Minute)

	return &MySQL{db}
}

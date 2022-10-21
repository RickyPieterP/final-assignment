package config

import (
	"fmt"

	"github.com/gomodul/envy"
)

type databaseItem struct {
	DriverName     string
	DataSourceName string
}

type database struct {
	MySQL databaseItem
}

var DB = database{
	MySQL: databaseItem{
		DriverName: envy.Get("MYSQL_DRIVER", "mysql"),
		DataSourceName: fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			envy.Get("MYSQL_USERNAME"),
			envy.Get("MYSQL_PASSWORD"),
			envy.Get("MYSQL_HOST"),
			envy.Get("MYSQL_PORT"),
			envy.Get("MYSQL_Name", "mygram"),
		),
	},
}

package database

import (
	"github.com/scalent-sushil/user-management-go/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Connect this function is use to connect the database.
func Connect() (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(config.DBURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

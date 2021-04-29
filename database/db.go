package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//Connect this function is use to connect the database. (*gorm.DB, error)
func Connect() {

	db, _ := gorm.Open(mysql.Open("root:Suchil$123@tcp(localhost:3306)/Golang?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	// if err != nil {
	// 	return nil, err
	// }

	sqldb, _ := db.DB()
	// if err != nil {
	// 	return nil, err
	// }
	sqldb.SetMaxIdleConns(50)

	sqldb.SetMaxOpenConns(100)

	DB = db
	// return DB, nil
}

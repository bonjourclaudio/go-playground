package db

import (
	"github.com/claudioontheweb/go-playground/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {

	config.GetConfig()

	// Get DB config from config.json
	var dbHost string = viper.GetString("DB_HOST")
	var dbUser string = viper.GetString("DB_USERNAME")
	var dbName string = viper.GetString("DB_NAME")
	var dbPassword string = viper.GetString("DB_PASSWORD")

	// Connect without DB
	db, dbError := gorm.Open("mysql", dbUser+":"+ dbPassword +"@tcp(" + dbHost+ ":3306)/")
	if dbError != nil {
		panic(dbError)
	}

	defer db.Close()

	// Create DB
	db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	db.Close()

	// Connect to DB
	db, dbError2 := gorm.Open("mysql", dbUser+":"+ dbPassword +"@tcp(" + dbHost+ ":3306)/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	if dbError2 != nil {
		panic(dbError2)
	}

	db.DB().SetMaxIdleConns(0)

	return db

}
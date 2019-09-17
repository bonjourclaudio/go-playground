package db

import (
	"github.com/claudioontheweb/gorm-rest-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {

	config.GetConfig()

	// Get DB config from config.json
	var dbHost string = viper.GetString("DB_HOST")
	var dbName string = viper.GetString("DB_NAME")
	var dbUser string = viper.GetString("DB_USERNAME")
	var dbPassword string = viper.GetString("DB_PASSWORD")

	// Connect to DB
	db, dbError := gorm.Open("mysql", dbUser+":"+ dbPassword +"@tcp(" + dbHost+ ":3306)/"+ dbName + "?charset=utf8&parseTime=True&loc=Local")
	if dbError != nil {
		panic("Failed to connect to database")
	}
	db.DB().SetMaxIdleConns(0)

	return db

}
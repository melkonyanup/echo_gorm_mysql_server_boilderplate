package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitDatabase() (*gorm.DB, error) {
	dbUserAndPassword := viper.GetString("mysql.username") + ":" + os.Getenv("DB_USER_PASSWORD")
	dbHostAndPort := viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port")
	dbName := viper.GetString("mysql.dbname")

	db, err := gorm.Open(mysql.Open(dbUserAndPassword+"@tcp("+dbHostAndPort+")/"+dbName+
		"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

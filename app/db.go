package app

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var dbConnection *gorm.DB

func SetupDatabase() {

	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	name := viper.GetString("database.dbname")
	host := viper.GetString("database.hostname")
	maxIdleConnection := viper.GetInt("database.maxIdleConnection")
	maxOpenConnection := viper.GetInt("database.maxOpenConnection")

	dbUri := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, name)
	//log.Println(dbUri)
	log.Println("Database connected...")

	conn, err := gorm.Open("mysql", dbUri)

	if err != nil {
		log.Fatal(err)
	}

	dbConnection = conn

	//log.Println(dbConnection.DB().Ping())

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	dbConnection.DB().SetMaxIdleConns(maxIdleConnection)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	dbConnection.DB().SetMaxOpenConns(maxOpenConnection)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	dbConnection.DB().SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	return dbConnection
}

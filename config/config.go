package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/fazriachyar/cloudBread/libraries"
	"github.com/spf13/viper"
)

func init(){
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
}

func GetString(key string)(string) {
	return viper.GetString(key)
}

func GetInt(key string)(int) {
	return viper.GetInt(key)
}

// const (
// 	DB_USER = "fazriachyar11"
// 	DB_PASSWORD = "200920"
// 	DB_NAME = "cloudBread"
// )

func SetupDB() *sql.DB {
	//dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	//os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) //dbinfo

	libraries.CheckErr(err)

	return db
}
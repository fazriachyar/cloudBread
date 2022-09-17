package config

import (
	"database/sql"
	"fmt"
	_ "os"

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

func SetupDB() *sql.DB {
	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", GetString("db.user"), GetString("db.pwd"), GetString("db.name"))
	//os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbinfo) //dbinfo
	db.Exec("CREATE TABLE IF NOT EXISTS bread (id SERIAL,breadid VARCHAR(50) NOT NULL, breadname VARCHAR(50) NOT NULL, price VARCHAR(50) NOT NULL, imgurl VARCHAR(50), PRIMARY KEY (id))");

	libraries.CheckErr(err)

	return db
}
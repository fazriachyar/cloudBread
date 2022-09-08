package models

import (
	"database/sql"
	_ "fmt"
	"os"

	"github.com/fazriachyar/cloudBread/libraries"
)

type JsonResponse struct {
	Type string `json:"type"`
	Data []Bread `json:"data"`
	Message string `json:"message"`
}

type Bread struct {
	BreadID string `json:"breadid"`
	BreadName string `json:"breadname"`
	BreadPrice string `json:"breadprice"`
	ImgURL string `json:"imgurl"`
}

// const (
// 	DB_USER = "fazriachyar11"
// 	DB_PASSWORD = "200920"
// 	DB_NAME = "cloudBread"
// )

func SetupDB() *sql.DB {
	//dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) 

	libraries.CheckErr(err)

	return db
}
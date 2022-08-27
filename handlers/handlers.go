package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Bread struct {
	BreadID string `json:"breadid"`
	BreadName string `json:"breadname"`
	BreadPrice string `json:"breadprice`
	ImgURL string `json:"imgurl"`
}

type JsonResponse struct {
	Type string `json:"type"`
	Data []Bread `json:"data"`
	Message string `json:"message"`
}

const (
	DB_USER = "fazriachyar11"
	DB_PASSWORD = "200920"
	DB_NAME = "cloudBread"
)

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

func GetBreads(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Sedang memuat Cloud Bread ...")

	//get all bread from bread table
	rows, err := db.Query("SELECT * FROM bread")

	//cek error
	checkErr(err)

	//var response []JsonResponse
	var breads []Bread

	//foreach bread
	for rows.Next() {
		var id int
		var BreadID string
		var BreadName string
		var BreadPrice string
		var ImgURL string

		err = rows.Scan(&id, &BreadID, &BreadName, &BreadPrice, &ImgURL)

		//cek err
		checkErr(err)

		breads = append(breads, Bread{BreadID: BreadID, BreadName: BreadName, BreadPrice: BreadPrice, ImgURL: ImgURL})
	}

	var response = JsonResponse{Type: "success", Data: breads}

	json.NewEncoder(w).Encode(response)
}

func GetBread(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	breadID := params["breadid"]

	var response = JsonResponse{}

	if breadID == "" {
		response = JsonResponse{Type: "error", Message: "Please insert ID bread..."}
	} else {
		db := setupDB()

		printMessage("Showing Bread by id")

		rows, err := db.Query("SELECT * FROM bread WHERE breadid = $1", breadID)
		checkErr(err)
		var breads []Bread

		for rows.Next() {
			var id int
			var BreadID string
			var BreadName string
			var BreadPrice string
			var ImgURL string
	
			err = rows.Scan(&id, &BreadID, &BreadName, &BreadPrice, &ImgURL)
	
			//cek err
			checkErr(err)
	
			breads = append(breads, Bread{BreadID: BreadID, BreadName: BreadName, BreadPrice: BreadPrice, ImgURL: ImgURL})
		}
		

		response = JsonResponse{Type: "success", Data: breads}
	
	}
	json.NewEncoder(w).Encode(response)
}

func CreateBread(w http.ResponseWriter, r *http.Request) {
	breadID := r.FormValue("breadid")
	breadName := r.FormValue("breadname")
	breadPrice := r.FormValue("breadprice")
	ImgURL := r.FormValue("imgurl")

	var response = JsonResponse{}

	if breadID == "" || breadName == "" || breadPrice == "" || ImgURL == "" {
		response = JsonResponse{Type: "error", Message: "Please Insert data..."}
	} else {
		db := setupDB()

		printMessage("Making new Bread...")

		fmt.Println("Making new Bread with ID: " + breadID + " and name: " + breadName)

		var lastInsertID int
		
		err := db.QueryRow("INSERT INTO bread(breadid, breadname, price, imgurl) VALUES($1, $2, $3, $4) returning id;", breadID, breadName, breadPrice, ImgURL).Scan(&lastInsertID)

		checkErr(err)

		response = JsonResponse{Type: "success", Message: "Bread has been made successfully!"}

	}

	json.NewEncoder(w).Encode(response)
}

func DeleteBread(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	breadID := params["breadid"]

	var response = JsonResponse{}

	if breadID == "" {
		response = JsonResponse{Type: "error", Message: "Please insert ID bread..."}
	} else {
		db := setupDB()

		printMessage("Deleting Bread...")

		_, err := db.Exec("DELETE FROM bread where breadid = $1", breadID)

		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The Bread has been deleted..."}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteBreads(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Deleting all breads...")

	_, err := db.Exec("DELETE FROM breads")

	checkErr(err)

	printMessage("All breads have been deleted !")

	var response = JsonResponse{Type: "success", Message: "All breads have been deleted!"}

	json.NewEncoder(w).Encode(response)
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

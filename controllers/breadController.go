package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"os"

	"github.com/fazriachyar/cloudBread/config"
	"github.com/fazriachyar/cloudBread/libraries"
	"github.com/fazriachyar/cloudBread/models"
	"github.com/fazriachyar/cloudBread/models/response"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)




func GetBreadsEndpoint(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	printMessage("Sedang memuat Cloud Bread ...")
	db.Exec("CREATE TABLE IF NOT EXISTS bread (id SERIAL,breadid VARCHAR(50) NOT NULL, breadname VARCHAR(50) NOT NULL, breadprice VARCHAR(50) NOT NULL, imgurl VARCHAR(50), PRIMARY KEY (id))");
	//get all bread from bread table
	rows, err := db.Query("SELECT * FROM bread")

	//cek error
	libraries.CheckErr(err)

	//var response []JsonResponse
	var breads []models.Bread

	//foreach bread
	for rows.Next() {
		var id int
		var BreadID string
		var BreadName string
		var BreadPrice string
		var ImgURL string

		err = rows.Scan(&id, &BreadID, &BreadName, &BreadPrice, &ImgURL)

		//cek err
		libraries.CheckErr(err)

		breads = append(breads, models.Bread{BreadID: BreadID, BreadName: BreadName, BreadPrice: BreadPrice, ImgURL: ImgURL})
	}

	var response = response.JsonResponse{Type: "success", Data: breads}

	json.NewEncoder(w).Encode(response)
}

func GetBreadEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	breadID := params["breadid"]

	var response = response.JsonResponse{}

	if breadID == "" {
		response.Type = "error"
		response.Message = "Please insert ID bread..."
	} else {
		db := config.SetupDB()
		db.Exec("CREATE TABLE IF NOT EXISTS bread (id SERIAL,breadid VARCHAR(50) NOT NULL, breadname VARCHAR(50) NOT NULL, price VARCHAR(50) NOT NULL, imgurl VARCHAR(50), PRIMARY KEY (id))");
		printMessage("Showing Bread by id")

		rows, err := db.Query("SELECT * FROM bread WHERE breadid = $1", breadID)
		libraries.CheckErr(err)
		var breads []models.Bread

		for rows.Next() {
			var id int
			var BreadID string
			var BreadName string
			var BreadPrice string
			var ImgURL string
	
			err = rows.Scan(&id, &BreadID, &BreadName, &BreadPrice, &ImgURL)
	
			//cek err
			libraries.CheckErr(err)
	
			breads = append(breads, models.Bread{BreadID: BreadID, BreadName: BreadName, BreadPrice: BreadPrice, ImgURL: ImgURL})
		}
		// response = models.JsonResponse{Type: "success", Data: breads}
		response.Type = "success"
		response.Data = breads
	}
	json.NewEncoder(w).Encode(response)
}

func CreateBreadEndpoint(w http.ResponseWriter, r *http.Request) {
	var response response.JsonResponse

	breadID := r.FormValue("breadid")
	breadName := r.FormValue("breadname")
	breadPrice := r.FormValue("breadprice")
	ImgURL := r.FormValue("imgurl")

	if breadID == "" || breadName == "" || breadPrice == "" || ImgURL == "" {
		// response = models.JsonResponse{Type: "error", Message: "Please Insert data..."}
		response.Type = "error"
		response.Message = "Please Insert data..."
	} else {
		db := config.SetupDB()
		
		db.Exec("CREATE TABLE IF NOT EXISTS bread (id SERIAL,breadid VARCHAR(50) NOT NULL, breadname VARCHAR(50) NOT NULL, price VARCHAR(50) NOT NULL, imgurl VARCHAR(50), PRIMARY KEY (id))");
		printMessage("Making new Bread...")

		fmt.Println("Making new Bread with ID: " + breadID + " and name: " + breadName)

		var lastInsertID int
		
		err := db.QueryRow("INSERT INTO bread(breadid, breadname, price, imgurl) VALUES($1, $2, $3, $4) returning id;", breadID, breadName, breadPrice, ImgURL).Scan(&lastInsertID)

		libraries.CheckErr(err)

		// response = models.JsonResponse{Type: "success", Message: "Bread has been made successfully!"}
		response.Type = "success"
		response.Message = "Bread has been made successfully!"
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteBreadEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	breadID := params["breadid"]

	var response response.JsonResponse
	
	if breadID == "" {
		// response = response.JsonResponse{Type: "error", Message: "Please insert ID bread..."}
		response.Type = "error"
		response.Message = "Please insert ID bread..."
		fmt.Printf(response.Message)
	} else {
		
		db := config.SetupDB()
		db.Exec("CREATE TABLE IF NOT EXISTS bread (id SERIAL,breadid VARCHAR(50) NOT NULL, breadname VARCHAR(50) NOT NULL, price VARCHAR(50) NOT NULL, imgurl VARCHAR(50), PRIMARY KEY (id))");
		printMessage("Deleting Bread...")

		_, err := db.Exec("DELETE FROM bread where breadid = $1", breadID)

		libraries.CheckErr(err)

		// response = models.JsonResponse{Type: "success", Message: "The Bread has been deleted..."}
		response.Type = "success"
		response.Message = "The Bread has been deleted..."
		fmt.Printf(response.Message)
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteBreadsEndpoint(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()
	db.Exec("CREATE TABLE IF NOT EXISTS bread (id SERIAL,breadid VARCHAR(50) NOT NULL, breadname VARCHAR(50) NOT NULL, price VARCHAR(50) NOT NULL, imgurl VARCHAR(50), PRIMARY KEY (id))");	
	printMessage("Deleting all breads...")

	_, err := db.Exec("DELETE FROM breads")

	libraries.CheckErr(err)

	printMessage("All breads have been deleted !")

	var response = response.JsonResponse{Type: "success", Message: "All breads have been deleted!"}
	fmt.Printf(response.Message)
	json.NewEncoder(w).Encode(response)
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

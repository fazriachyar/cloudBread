package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	libraries.PrintMessage("Sedang memuat Cloud Bread ...")
	rows, err := db.Query("SELECT * FROM bread")

	libraries.CheckErr(err)

	var breads []models.Bread
	for rows.Next() {
		var (
			id int
			BreadID string
			BreadName string
			BreadPrice string
			ImgURL string
		)

		err = rows.Scan(&id, &BreadID, &BreadName, &BreadPrice, &ImgURL)
		libraries.CheckErr(err)

		breads = append(breads, models.Bread{BreadID: BreadID, BreadName: BreadName, BreadPrice: BreadPrice, ImgURL: ImgURL})
	}
	var response = response.JsonResponse{Type: "success", Data: breads}

	json.NewEncoder(w).Encode(response)
}

func GetBreadEndpoint(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		breadID = params["breadid"]
		response = response.JsonResponse{}
	)

	if breadID == "" {
		response.Type = "error"
		response.Message = "Please insert ID bread..."
	} else {
		var (
			db = config.SetupDB()
			breads []models.Bread
		)
		libraries.PrintMessage("Showing Bread by id")

		rows, err := db.Query("SELECT * FROM bread WHERE breadid = $1", breadID)
		libraries.CheckErr(err)

		for rows.Next() {
			var (
				id int
				BreadID string
				BreadName string
				BreadPrice string
				ImgURL string
			)
	
			err = rows.Scan(&id, &BreadID, &BreadName, &BreadPrice, &ImgURL)
	
			//cek err
			libraries.CheckErr(err)
	
			breads = append(breads, models.Bread{BreadID: BreadID, BreadName: BreadName, BreadPrice: BreadPrice, ImgURL: ImgURL})
		}
		response.Type = "success"
		response.Data = breads
	}
	json.NewEncoder(w).Encode(response)
}

func CreateBreadEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		response response.JsonResponse
		bread models.Bread
	)

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &bread)

	if bread.BreadID == "" || bread.BreadName == "" || bread.BreadPrice == "" || bread.ImgURL == "" {
		response.Type = "error"
		response.Message = "Please Insert data..."

	} else {
		db := config.SetupDB()
		
		libraries.PrintMessage("Making new Bread...")
		fmt.Println("Making new Bread with ID: " + bread.BreadID + " and name: " + bread.BreadName)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO bread(breadid, breadname, price, imgurl) VALUES($1, $2, $3, $4) returning id;", bread.BreadID, bread.BreadName, bread.BreadPrice, bread.ImgURL).Scan(&lastInsertID)

		libraries.CheckErr(err)

		response.Type = "success"
		response.Message = "Bread has been made successfully!"
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateBreadEndpoint(w http.ResponseWriter, r *http.Request) {
	var (
		response response.JsonResponse
		bread models.Bread
		params = mux.Vars(r)
		breadIDParam = params["breadid"]	
	)

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &bread)

	if breadIDParam == "" {
		response.Type = "error"
		response.Message = "Please insert ID bread..."
		fmt.Printf(response.Message)
	} else {
		db := config.SetupDB()
		libraries.PrintMessage("Updating Bread...")

		_, err := db.Exec("UPDATE bread SET breadid = $1, breadname= $2, price = $3, imgurl = $4 WHERE breadid = $5;", bread.BreadID, bread.BreadName, bread.BreadPrice, bread.ImgURL, breadIDParam)
		libraries.CheckErr(err)

		response.Type = "success"
		response.Message = "Bread with ID : " + breadIDParam + " has been updated"
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteBreadEndpoint(w http.ResponseWriter, r *http.Request) {
	var (
		response response.JsonResponse
		params = mux.Vars(r)
		breadID = params["breadid"]
	)
	
	if breadID == "" {
		response.Type = "error"
		response.Message = "Please insert ID bread..."
		fmt.Printf(response.Message)

	} else {	
		db := config.SetupDB()
		libraries.PrintMessage("Deleting Bread...")
		_, err := db.Exec("DELETE FROM bread where breadid = $1", breadID)

		libraries.CheckErr(err)

		response.Type = "success"
		response.Message = "The Bread has been deleted..."
		fmt.Printf(response.Message)
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteBreadsEndpoint(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()
	libraries.PrintMessage("Deleting all breads...")
	_, err := db.Exec("DELETE FROM breads")

	libraries.CheckErr(err)
	libraries.PrintMessage("All breads have been deleted !")

	var response = response.JsonResponse{Type: "success", Message: "All breads have been deleted!"}
	fmt.Printf(response.Message)
	json.NewEncoder(w).Encode(response)
}

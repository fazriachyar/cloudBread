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

	printMessage("Sedang memuat Cloud Bread ...")
	db.Exec("CREATE TABLE IF NOT EXISTS bread (id SERIAL,breadid VARCHAR(50) NOT NULL, breadname VARCHAR(50) NOT NULL, price VARCHAR(50) NOT NULL, imgurl VARCHAR(50), PRIMARY KEY (id))");
	//get all bread from bread table
	rows, err := db.Query("SELECT * FROM bread")

	//cek error
	libraries.CheckErr(err)

	//var response []JsonResponse
	var breads []models.Bread

	//foreach bread
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
		db.Exec("CREATE TABLE IF NOT EXISTS bread (id SERIAL,breadid VARCHAR(50) NOT NULL, breadname VARCHAR(50) NOT NULL, price VARCHAR(50) NOT NULL, imgurl VARCHAR(50), PRIMARY KEY (id))");
		printMessage("Showing Bread by id")

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
		// response = models.JsonResponse{Type: "success", Data: breads}
		response.Type = "success"
		response.Data = breads
	}
	json.NewEncoder(w).Encode(response)
}

func CreateBreadEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response response.JsonResponse
	var bread models.Bread

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &bread)

	// breadID := r.FormValue("breadid")
	// breadName := r.FormValue("breadname")
	// breadPrice := r.FormValue("breadprice")
	// ImgURL := r.FormValue("imgurl")

	if bread.BreadID == "" || bread.BreadName == "" || bread.BreadPrice == "" || bread.ImgURL == "" {
		// response = models.JsonResponse{Type: "error", Message: "Please Insert data..."}
		response.Type = "error"
		response.Message = "Please Insert data..."
	} else {
		db := config.SetupDB()
		
		db.Exec("CREATE TABLE IF NOT EXISTS bread (id SERIAL,breadid VARCHAR(50) NOT NULL, breadname VARCHAR(50) NOT NULL, price VARCHAR(50) NOT NULL, imgurl VARCHAR(50), PRIMARY KEY (id))");
		printMessage("Making new Bread...")

		fmt.Println("Making new Bread with ID: " + bread.BreadID + " and name: " + bread.BreadName)

		var lastInsertID int
		
		err := db.QueryRow("INSERT INTO bread(breadid, breadname, price, imgurl) VALUES($1, $2, $3, $4) returning id;", bread.BreadID, bread.BreadName, bread.BreadPrice, bread.ImgURL).Scan(&lastInsertID)

		libraries.CheckErr(err)

		// response = models.JsonResponse{Type: "success", Message: "Bread has been made successfully!"}

		response.Type = "success"
		response.Message = "Bread has been made successfully!"
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateBreadEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	breadIDParam := params["breadid"]

	var response response.JsonResponse
	var bread models.Bread

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &bread)

	// breadID := r.FormValue("breadid")
	// breadName := r.FormValue("breadname")
	// breadPrice := r.FormValue("breadprice")
	// ImgURL := r.FormValue("imgurl")

	if breadIDParam == "" {
		response.Type = "error"
		response.Message = "Please insert ID bread..."
		fmt.Printf(response.Message)
	} else {
		db := config.SetupDB()
		db.Exec("CREATE TABLE IF NOT EXISTS bread (id SERIAL,breadid VARCHAR(50) NOT NULL, breadname VARCHAR(50) NOT NULL, price VARCHAR(50) NOT NULL, imgurl VARCHAR(50), PRIMARY KEY (id))");
		printMessage("Updating Bread...")

		_, err := db.Exec("UPDATE bread SET breadid = $1, breadname= $2, price = $3, imgurl = $4 WHERE breadid = $5;", bread.BreadID, bread.BreadName, bread.BreadPrice, bread.ImgURL, breadIDParam)
		libraries.CheckErr(err)

		response.Type = "success"
		response.Message = "Bread with ID : " + breadIDParam + " has been updated"
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

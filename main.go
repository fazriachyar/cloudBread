package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fazriachyar/cloudBread/config"
	"github.com/fazriachyar/cloudBread/controllers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)



func main() {
	// port := os.Getenv("PORT")
	// if port == "" {
    //     log.Fatal("$PORT must be set")
    // }
	
	router := mux.NewRouter()

	//GET ALL BREAD
	router.HandleFunc("/breads/", controllers.GetBreadsEndpoint).Methods("GET")

	//GET BREAD by id
	router.HandleFunc("/breads/{breadid}", controllers.GetBreadEndpoint).Methods("GET")

	//CREATE BREAD
	router.HandleFunc("/breads/", controllers.CreateBreadEndpoint).Methods("POST")

	//DELETE BREAD by id
	router.HandleFunc("/breads/{breadid}", controllers.DeleteBreadEndpoint).Methods("DELETE")

	//DELETE ALL BREAD
	router.HandleFunc("/breads/", controllers.DeleteBreadsEndpoint).Methods("DELETE")


	//serve
	fmt.Println("Server started at :1337")
	log.Fatal(http.ListenAndServe(config.GetString("server.address"), router))

}

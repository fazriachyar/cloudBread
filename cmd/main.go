package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fazriachyar/cloudBread/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)



func main() {
	router := mux.NewRouter()

	//GET ALL BREAD
	router.HandleFunc("/breads/", handlers.GetBreads).Methods("GET")

	//GET BREAD by id
	router.HandleFunc("/breads/{breadid}", handlers.GetBread).Methods("GET")

	//CREATE BREAD
	router.HandleFunc("/breads/", handlers.CreateBread).Methods("POST")

	//DELETE BREAD by id
	router.HandleFunc("/breads/{breadid}", handlers.DeleteBread).Methods("DELETE")

	//DELETE ALL BREAD
	router.HandleFunc("/breads/", handlers.DeleteBreads).Methods("DELETE")

	//serve
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}




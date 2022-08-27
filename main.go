package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fazriachyar/cloudBread/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)



func main() {
	port := os.Getenv("PORT")
	if port == "" {
        log.Fatal("$PORT must be set")
    }
	
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
	fmt.Println("Server at port")
	log.Fatal(http.ListenAndServe(":" + port, router))

}

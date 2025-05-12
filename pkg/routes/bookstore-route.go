package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/sumanth/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){

	fmt.Println("Enterd Routes")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
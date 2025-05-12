package main

// this is first intilized and then it tells the golang where the routes are
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sumanth/go-bookstore/pkg/routes"
)

func main(){
	fmt.Println("Enterd main")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	//  Listen and serve is a function which helps to create a server
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	

}
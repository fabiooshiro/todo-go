package main

import (
	"curso/aula-07-mvc/controller"
	"curso/aula-07-mvc/model"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // mysql
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

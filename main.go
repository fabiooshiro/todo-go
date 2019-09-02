package main

import (
	"curso/aula-07-mvc/controller"
	"curso/aula-07-mvc/model"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sanyogpatel-tecblic/To-Do/controller/config"
	"github.com/sanyogpatel-tecblic/To-Do/controller/routes"
	// "github.com/sanyogpatel-tecblic/To-Do/controller"
)

var app config.AppConfig

func main() {
	fmt.Println("Server is getting started...")
	fmt.Println("Listening at port 8010 ...")
	http.ListenAndServe(":8010", routes.Routes(&app))

	routes.Routes(&app)
}

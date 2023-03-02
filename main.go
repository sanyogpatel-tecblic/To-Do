package main

import (
	_ "github.com/lib/pq"
	"github.com/sanyogpatel-tecblic/To-Do/controller/routes"
	// "github.com/sanyogpatel-tecblic/To-Do/controller"
)

func main() {
	routes.Router()
}
